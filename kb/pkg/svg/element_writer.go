package svg

import (
	"fmt"
	"kb/pkg/models"
	"log"
	"strings"
)

type elementWriter struct {
	options *Options
}

func newElementWriter(options *Options) *elementWriter {
	return &elementWriter{options: options}
}

func (ew *elementWriter) writeElement(w *xmlWriter, element models.KeyboardElement) {
	if !element.GetVisible() {
		return
	}

	w.writeStartElement("g")

	ew.writeAttributes(w, element)
	ew.writeSubElements(w, element)

	w.writeEndElement()
}

func (ew *elementWriter) writeAttributes(w *xmlWriter, element models.KeyboardElement) {
	w.writeAttributeString("id", element.GetName())
	ew.writeTransform(w, element)
}

func (ew *elementWriter) writeSubElements(w *xmlWriter, element models.KeyboardElement) {
	if element.GetDebug() {
		ew.writeDebugOverlay(w, element)
	}
}

func (ew *elementWriter) writeDebugOverlay(w *xmlWriter, element models.KeyboardElement) {
	// Write border (With Margin)
	{
		height := element.GetHeight() + (element.GetMargin() * 2)
		width := element.GetWidth() + (element.GetMargin() * 2)
		w.writeStartElement("path")
		w.writeAttributeString("id", fmt.Sprintf("%sDebugOverlayMargin", element.GetName()))
		w.writeAttributeString("d", fmt.Sprintf("M -%f,-%f h %f v %f h -%f v -%f h %f",
			width/2, height/2, width, height, width, height, width))
		w.writeAttributeString("style", "fill:none;stroke:#00ff00;stroke-width:0.1")
		w.writeEndElement()
	}

	// Write border (Without Margin)
	{
		height := element.GetHeight()
		width := element.GetWidth()
		w.writeStartElement("path")
		w.writeAttributeString("id", fmt.Sprintf("%sDebugOverlayBorder", element.GetName()))
		w.writeAttributeString("d", fmt.Sprintf("M -%f,-%f h %f v %f h -%f v -%f h %f",
			width/2, height/2, width, height, width, height, width))
		w.writeAttributeString("style", "fill:none;stroke:#00ffff;stroke-width:0.1")
		w.writeEndElement()
	}

	// Write center axes
	{
		w.writeStartElement("path")
		w.writeAttributeString("id", fmt.Sprintf("%sCenterAxes", element.GetName()))
		w.writeAttributeString("d", "m 0,0 h 5 h -10 h 5 v 5 v -10 v 5")
		w.writeAttributeString("style", "fill:none;stroke:#00ffff;stroke-width:0.1")
		w.writeEndElement()
	}
}

func (ew *elementWriter) writeTransform(w *xmlWriter, element models.KeyboardElement) {
	transformStrings := []string{}

	xOffset := element.GetXOffset()
	yOffset := element.GetYOffset()

	stack, ok := element.GetParent().(*models.Stack)
	if ok {
		log.Print("elementWriter::writeTransform() - parent is stack!")
		if stack.Orientation == models.StackOrientationHorizontal {
			xOffset += getOffsetInStack(stack, element)
		} else if stack.Orientation == models.StackOrientationVertical {
			yOffset += getOffsetInStack(stack, element)
		} else {
			panic("unknown stack orientation")
		}
	}

	if xOffset != 0.0 || yOffset != 0.0 {
		transformStrings = append(transformStrings, fmt.Sprintf("translate(%f,%f)", xOffset, yOffset))
	}

	if element.GetRotation() != 0.0 {
		transformStrings = append(transformStrings, fmt.Sprintf("rotate(%f)", element.GetRotation()))
	}

	if len(transformStrings) > 0 {
		w.writeAttributeString("transform", strings.Join(transformStrings, " "))
	}
}
