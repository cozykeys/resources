package svg

import (
	"kb/pkg/models"
)

const (
	DefaultPathFill        = "none"
	DefaultPathFillOpacity = "1"
	DefaultPathStroke      = "#000000"
	DefaultPathStrokeWidth = "0.5"
)

type pathWriter struct {
	options *Options
}

func newPathWriter(options *Options) *pathWriter {
	return &pathWriter{options: options}
}

func (pw *pathWriter) write(w *xmlWriter, path *models.Path) {
	if !path.Visible {
		return
	}

	w.writeStartElement("g")

	_elementWriter := newElementWriter(pw.options)

	_elementWriter.writeAttributes(w, path)
	pw.writeAttributes(w, path)

	_elementWriter.writeSubElements(w, path)
	pw.writeSubElements(w, path)

	w.writeEndElement()
}

func (pw *pathWriter) writeAttributes(w *xmlWriter, path *models.Path) {
}

func (pw *pathWriter) writeSubElements(w *xmlWriter, path *models.Path) {
	w.writeStartElement("path")
	w.writeAttributeString("style", cssStyleString(map[string]string{
		"fill":         stringOrDefault(path.Fill, DefaultPathFill),
		"fill-opacity": stringOrDefault(path.FillOpacity, DefaultPathFillOpacity),
		"stroke":       stringOrDefault(path.Stroke, DefaultPathStroke),
		"stroke-width": stringOrDefault(path.StrokeWidth, DefaultPathStrokeWidth),
	}))
	w.writeAttributeString("d", path.Data())
	w.writeEndElement()
}
