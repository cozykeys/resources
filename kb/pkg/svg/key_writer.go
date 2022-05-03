package svg

import (
	"fmt"
	"kb/pkg/models"
	"log"
	"strings"

	"github.com/beevik/etree"
)

func writeKey(parent *etree.Element, key *models.Key) error {
	writer := &keyWriter{
		parent: parent,
		key:    key,
	}

	return writer.write()
}

type keyWriter struct {
	parent *etree.Element
	key    *models.Key
}

func (w *keyWriter) write() error {
	if !w.key.Visible {
		return nil
	}

	g := w.parent.CreateElement("g")

	// Attributes
	g.CreateAttr("id", w.key.Name)
	writeTransform(g, w.key)

	// Elements
	if w.key.Debug {
		log.Printf("Writing debug overlay...")
		writeDebugOverlay(g, w.key)
	}

	w.writeSwitchCutoutPath(g)
	w.writeOverlayV1(g)
	//w.writeOverlayV2(g)
	w.writeLegends(g)

	return nil
}

func (w *keyWriter) writeSwitchCutoutPath(parent *etree.Element) {
	e := parent.CreateElement("path")
	e.CreateAttr("style", "fill:none;fill-opacity:1;stroke:#000000;stroke-width:0.5")
	e.CreateAttr("d", switchCutoutPathData)
}

func (w *keyWriter) writeOverlayV1(parent *etree.Element) error {
	if options == nil || !options.KeycapOverlaysEnabled {
		return nil
	}

	e := parent.CreateElement("path")
	e.CreateAttr("id", fmt.Sprintf("%sKeycapOverlayInner", w.key.Name))
	e.CreateAttr("style", cssStyleString(map[string]string{
		"fill":         stringOrDefault(w.key.Fill, "#ffffff"),
		"stroke":       stringOrDefault(w.key.Stroke, "#000000"),
		"stroke-width": "0.5",
	}))
	e.CreateAttr("d", fmt.Sprintf("M -%f,-%f h %f v %f h -%f v -%f h %f",
		w.key.Width/2,
		w.key.Height/2,
		w.key.Width,
		w.key.Height,
		w.key.Width,
		w.key.Height,
		w.key.Width,
	))

	/*
	   // Next we write it with a style that is more visually pleasing
	   writer.WriteStartElement("path");
	   writer.WriteAttributeString("id", $"{key.Name}KeycapOverlay");
	   writer.WriteAttributeString("d", $"M -{w / 2},-{h / 2} h {w} v {h} h -{w} v -{h} h {w}");

	   var styleDictionary = new Dictionary<string, string>
	   {
	       { "fill", !string.IsNullOrWhiteSpace(key.Fill) ? key.Fill : "#ffffff" },
	       { "stroke", !string.IsNullOrWhiteSpace(key.Stroke) ? key.Stroke : "#000000" },
	       { "stroke-width", "0.5" },
	   };

	   writer.WriteAttributeString("style", styleDictionary.ToCssStyleString());
	   writer.WriteEndElement();
	*/
	return nil
}

func (w *keyWriter) writeOverlayV2(parent *etree.Element) error {
	if options == nil || !options.KeycapOverlaysEnabled {
		return nil
	}

	{
		fill, err := parseColor(stringOrDefault(w.key.Fill, "#ffffff"))
		if err != nil {
			return err
		}

		// Make the outer overlay slightly darker
		fill.R = fill.R - (fill.R / 5)
		fill.G = fill.G - (fill.G / 5)
		fill.B = fill.B - (fill.B / 5)

		e := parent.CreateElement("path")
		e.CreateAttr("id", fmt.Sprintf("%sKeycapOverlayOuter", w.key.Name))
		e.CreateAttr("d", keycapPathDataOuter)
		e.CreateAttr("style", cssStyleString(map[string]string{
			"fill":         formatColorHex(fill),
			"stroke":       stringOrDefault(w.key.Stroke, "#000000"),
			"stroke-width": "0.5",
		}))
	}

	{
		e := parent.CreateElement("path")
		e.CreateAttr("id", fmt.Sprintf("%sKeycapOverlayInner", w.key.Name))
		e.CreateAttr("d", keycapPathDataInner)
		e.CreateAttr("style", cssStyleString(map[string]string{
			"fill":         stringOrDefault(w.key.Fill, "#ffffff"),
			"stroke":       stringOrDefault(w.key.Stroke, "#000000"),
			"stroke-width": "0.5",
		}))
	}

	return nil
}

func (w *keyWriter) writeLegends(parent *etree.Element) error {
	if options == nil || !options.KeycapOverlaysEnabled {
		return nil
	}

	legendIndex := 0
	for _, legend := range w.key.Legends {
		legend.Name = fmt.Sprintf("%sLegend%d", w.key.Name, legendIndex)
		err := writeLegend(parent, &legend)
		if err != nil {
			return err
		}
		legendIndex++
	}

	return nil
}

var switchCutoutPathData string = reduceSpaces(strings.Join([]string{
	"M  0.0 -7.0",
	"L  7.0 -7.0",
	"L  7.0 -6.0",
	"L  7.8 -6.0",
	"L  7.8  6.0",
	"L  7.0  6.0",
	"L  7.0  7.0",
	"L -7.0  7.0",
	"L -7.0  6.0",
	"L -7.8  6.0",
	"L -7.8 -6.0",
	"L -7.0 -6.0",
	"L -7.0 -7.0",
	"L  0.0 -7.0",
}, " "))

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
