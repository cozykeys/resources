package svg

import (
	"kb/pkg/models"
	"log"
)

type stackWriter struct {
	options *Options
}

func newStackWriter(options *Options) *stackWriter {
	return &stackWriter{options: options}
}

func (sw *stackWriter) write(w *xmlWriter, stack *models.Stack) {
	if !stack.Visible {
		return
	}

	if stack.Orientation == models.StackOrientationVertical {
		log.Print("stackWriter::write() - vertical")
	} else {
		log.Print("stackWriter::write() - horizontal")
	}

	w.writeStartElement("g")

	_elementWriter := newElementWriter(sw.options)
	_groupWriter := newGroupWriter(sw.options)

	_elementWriter.writeAttributes(w, stack)
	_groupWriter.writeAttributes(w, stack)
	sw.writeAttributes(w, stack)

	_elementWriter.writeSubElements(w, stack)
	_groupWriter.writeSubElements(w, stack)
	sw.writeSubElements(w, stack)

	w.writeEndElement()
}

func (sw *stackWriter) writeAttributes(w *xmlWriter, stack *models.Stack) {
}

func (sw *stackWriter) writeSubElements(w *xmlWriter, stack *models.Stack) {
}
