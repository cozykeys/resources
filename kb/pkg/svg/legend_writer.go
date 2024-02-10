package svg

import (
	"fmt"
	"kb/pkg/models"
)

type legendWriter struct {
	options *Options
}

func newLegendWriter(options *Options) *legendWriter {
	return &legendWriter{options: options}
}

func (lw *legendWriter) write(w *xmlWriter, legend *models.Legend) {
	if !legend.Visible {
		return
	}

	w.writeStartElement("text")

	_elementWriter := newElementWriter(lw.options)

	_elementWriter.writeAttributes(w, legend)
	lw.writeAttributes(w, legend)

	_elementWriter.writeSubElements(w, legend)
	lw.writeSubElements(w, legend)

	w.writeEndElement()
}

func (lw *legendWriter) writeAttributes(w *xmlWriter, legend *models.Legend) {
	// TODO: Is this necessary when the same thing is in the style?
	w.writeAttributeString("text-anchor", "middle")

	fontSize := float64OrDefault(legend.FontSize, 4.0)

	w.writeAttributeString("style", cssStyleString(map[string]string{
		"fill":              stringOrDefault(legend.Color, "#000000"),
		"dominant-baseline": "central",
		"text-anchor":       "middle",
		"font-size":         fmt.Sprintf("%fpx", fontSize),
		"font-family":       "sans-serif",
		"font-weight":       "normal",
		"font-style":        "normal",
	}))
}

func (lw *legendWriter) writeSubElements(w *xmlWriter, legend *models.Legend) {
	w.writeText(legend.Text)
}
