package svg

import (
	"fmt"
	"kb/pkg/models"
	"strings"
)

type keyWriter struct {
	options *Options
}

func newKeyWriter(options *Options) *keyWriter {
	return &keyWriter{options: options}
}

func (kw *keyWriter) write(w *xmlWriter, key *models.Key) {
	if !key.GetVisible() {
		return
	}

	w.writeStartElement("g")

	_elementWriter := newElementWriter(kw.options)

	_elementWriter.writeAttributes(w, key)
	kw.writeAttributes(w, key)

	_elementWriter.writeSubElements(w, key)
	kw.writeSubElements(w, key)

	w.writeEndElement()
}

func (kw *keyWriter) writeAttributes(w *xmlWriter, key *models.Key) {
}

func (kw *keyWriter) writeSubElements(w *xmlWriter, key *models.Key) {
	kw.writeSwitchCutoutPath(w, key)
	kw.writeKeycapOverlay(w, key)
	//kw.writeOverlayV1(w, key)
	//w.writeOverlayV2(w, key)
	kw.writeLegends(w, key)
}

func (kw *keyWriter) writeSwitchCutoutPath(w *xmlWriter, key *models.Key) {
	_pathWriter := newPathWriter(kw.options)
	_pathWriter.write(w, switchPath)
}

func (kw *keyWriter) writeOverlayV1(w *xmlWriter, key *models.Key) {
	if !kw.options.EnableKeycapOverlays {
		return
	}

	// TODO: Use path writer
	w.writeStartElement("path")
	w.writeAttributeString("id", fmt.Sprintf("%sKeycapOverlayInner", key.GetName()))
	w.writeAttributeString("style", cssStyleString(map[string]string{
		"fill":         stringOrDefault(key.Fill, "#ffffff"),
		"stroke":       stringOrDefault(key.Stroke, "#000000"),
		"stroke-width": "0.5",
	}))
	w.writeAttributeString("d", fmt.Sprintf("M -%f,-%f h %f v %f h -%f v -%f h %f",
		key.GetWidth()/2,
		key.GetHeight()/2,
		key.GetWidth(),
		key.GetHeight(),
		key.GetWidth(),
		key.GetHeight(),
		key.GetWidth(),
	))
	w.writeEndElement()
}

func (kw *keyWriter) writeOverlayV2(w *xmlWriter, key *models.Key) {
	if !kw.options.EnableKeycapOverlays {
		return
	}

	{
		fill, err := parseColor(stringOrDefault(key.Fill, "#ffffff"))
		if err != nil {
			panic("failed to parse color")
		}

		// Make the outer overlay slightly darker
		fill.R = fill.R - (fill.R / 5)
		fill.G = fill.G - (fill.G / 5)
		fill.B = fill.B - (fill.B / 5)

		w.writeStartElement("path")
		w.writeAttributeString("id", fmt.Sprintf("%sKeycapOverlayOuter", key.GetName()))
		w.writeAttributeString("d", keycapPathDataOuter)
		w.writeAttributeString("style", cssStyleString(map[string]string{
			"fill":         formatColorHex(fill),
			"stroke":       stringOrDefault(key.Stroke, "#000000"),
			"stroke-width": "0.5",
		}))
		w.writeEndElement()
	}

	{
		w.writeStartElement("path")
		w.writeAttributeString("id", fmt.Sprintf("%sKeycapOverlayInner", key.GetName()))
		w.writeAttributeString("d", keycapPathDataInner)
		w.writeAttributeString("style", cssStyleString(map[string]string{
			"fill":         stringOrDefault(key.Fill, "#ffffff"),
			"stroke":       stringOrDefault(key.Stroke, "#000000"),
			"stroke-width": "0.5",
		}))
		w.writeEndElement()
	}
}

func (kw *keyWriter) writeKeycapOverlay(w *xmlWriter, key *models.Key) {
	if !kw.options.EnableKeycapOverlays {
		return
	}

	width := key.GetWidth()
	height := key.GetHeight()

	// TODO: Make a models.Path and use pathWriter instead
	w.writeStartElement("path")
	w.writeAttributeString("id", fmt.Sprintf("%sKeycapOverlay", key.GetName()))
	w.writeAttributeString("d", fmt.Sprintf("M -%f,-%f h %f v %f h -%f v -%f h %f", width/2, height/2, width, height, width, height, width))
	w.writeAttributeString("style", cssStyleString(map[string]string{
		"fill":         stringOrDefault(key.Fill, "#ffffff"),
		"stroke":       stringOrDefault(key.Stroke, "#000000"),
		"stroke-width": "0.5",
	}))
	w.writeEndElement()
}

func (kw *keyWriter) writeLegends(w *xmlWriter, key *models.Key) {
	if !kw.options.EnableKeycapOverlays {
		return
	}

	legendIndex := 0
	for _, legend := range key.Legends {
		legend.Name = fmt.Sprintf("%sLegend%d", key.GetName(), legendIndex)
		_legendWriter := newLegendWriter(kw.options)
		_legendWriter.write(w, &legend)
		legendIndex++
	}
}

var switchPath *models.Path = &models.Path{
	Components: []models.PathComponent{
		&models.AbsoluteMoveTo{EndPoint: &models.Point{X: 0.0, Y: -7.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: 7.0, Y: -7.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: 7.0, Y: -6.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: 7.8, Y: -6.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: 7.8, Y: 6.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: 7.0, Y: 6.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: 7.0, Y: 7.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: -7.0, Y: 7.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: -7.0, Y: 6.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: -7.8, Y: 6.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: -7.8, Y: -6.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: -7.0, Y: -6.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: -7.0, Y: -7.0}},
		&models.AbsoluteLineTo{EndPoint: &models.Point{X: 0.0, Y: -7.0}},
	},
}

var keycapPathDataOuter string = reduceSpaces(strings.Join([]string{
	"M  0.0  -9.05",
	"L  8.05 -9.05",
	"Q  9.05 -9.05",
	"   9.05 -8.05",
	"L  9.05  8.05",
	"Q  9.05  9.05",
	"   8.05  9.05",
	"L -8.05  9.05",
	"Q -9.05  9.05",
	"  -9.05  8.05",
	"L -9.05 -8.05",
	"Q -9.05 -9.05",
	"  -8.05 -9.05",
	"L  0.0  -9.05",
}, " "))

var keycapPathDataInner string = reduceSpaces(strings.Join([]string{
	"M  0.0  -8.05",
	"L  4.05 -8.05",
	"Q  6.05 -8.05",
	"   6.05 -6.05",
	"L  6.05  4.55",
	"Q  6.05  6.55",
	"   4.05  6.55",
	"L -4.05  6.55",
	"Q -6.05  6.55",
	"  -6.05  4.55",
	"L -6.05 -6.05",
	"Q -6.05 -8.05",
	"  -4.05 -8.05",
	"L  0.0  -8.05",
}, " "))
