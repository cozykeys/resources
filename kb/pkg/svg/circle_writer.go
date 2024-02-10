package svg

import (
	"fmt"
	"kb/pkg/models"
)

const (
	DefaultCircleFill        = "none"
	DefaultCircleStroke      = "#0000ff"
	DefaultCircleStrokeWidth = "0.01"
)

type circleWriter struct {
	options *Options
}

func newCircleWriter(options *Options) *circleWriter {
	return &circleWriter{options: options}
}

func (cw *circleWriter) write(w *xmlWriter, circle *models.Circle) {
	if !circle.Visible {
		return
	}

	w.writeStartElement("g")

	_elementWriter := newElementWriter(cw.options)

	_elementWriter.writeAttributes(w, circle)
	cw.writeAttributes(w, circle)

	_elementWriter.writeSubElements(w, circle)
	cw.writeSubElements(w, circle)

	w.writeEndElement()
}

func (cw *circleWriter) writeAttributes(w *xmlWriter, circle *models.Circle) {
}

func (cw *circleWriter) writeSubElements(w *xmlWriter, circle *models.Circle) {
	// First we write it with the style that Ponoko expects
	w.writeStartElement("circle")
	//w.writeAttributeString("id", "TODO")
	w.writeAttributeString("r", fmt.Sprintf("%f", circle.Size/2))
	w.writeAttributeString("style", "fill:none;stroke:#000000;stroke-width:0.5")
	w.writeEndElement()

	if cw.options.EnableVisualSwitchCutouts {
		w.writeStartElement("circle")
		//w.writeAttributeString("id", "TODO")
		w.writeAttributeString("style", cssStyleString(map[string]string{
			"fill":         stringOrDefault(circle.Fill, DefaultCircleFill),
			"stroke":       stringOrDefault(circle.Stroke, DefaultCircleStroke),
			"stroke-width": stringOrDefault(circle.StrokeWidth, DefaultCircleStrokeWidth),
		}))
		w.writeAttributeString("r", fmt.Sprintf("%f", circle.Size/2))
		w.writeEndElement()
	}
}
