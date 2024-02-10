package svg

import (
	"kb/pkg/models"
)

const (
	DefaultTextAnchor = "middle"
	DefaultFont       = "30px sans-serif"
	DefaultFill       = "#000000"
)

type textWriter struct {
	options *Options
}

func newTextWriter(options *Options) *textWriter {
	return &textWriter{options: options}
}

func (tw *textWriter) write(w *xmlWriter, text *models.Text) {
	if !text.Visible {
		return
	}

	w.writeStartElement("g")

	_elementWriter := newElementWriter(tw.options)

	_elementWriter.writeAttributes(w, text)
	tw.writeAttributes(w, text)

	_elementWriter.writeSubElements(w, text)
	tw.writeSubElements(w, text)

	w.writeEndElement()
}

func (tw *textWriter) writeAttributes(w *xmlWriter, text *models.Text) {
}

func (tw *textWriter) writeSubElements(w *xmlWriter, text *models.Text) {
	if !tw.options.EnableVisualSwitchCutouts {
		return
	}

	w.writeStartElement("text")
	w.writeAttributeString("text-anchor", stringOrDefault(text.TextAnchor, ""))
	w.writeAttributeString("style", cssStyleString(map[string]string{
		"font": stringOrDefault(text.Font, DefaultFont),
		"fill": stringOrDefault(text.Fill, DefaultFill),
	}))
	w.writeText(text.Content)
	w.writeEndElement()
}
