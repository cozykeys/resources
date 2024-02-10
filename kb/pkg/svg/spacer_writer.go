package svg

import "kb/pkg/models"

type spacerWriter struct {
	options *Options
}

func newSpacerWriter(options *Options) *spacerWriter {
	return &spacerWriter{options: options}
}

func (sw *spacerWriter) write(w *xmlWriter, spacer *models.Spacer) {
	if !spacer.Visible {
		return
	}

	w.writeStartElement("g")

	_elementWriter := newElementWriter(sw.options)

	_elementWriter.writeAttributes(w, spacer)
	sw.writeAttributes(w, spacer)

	_elementWriter.writeSubElements(w, spacer)
	sw.writeSubElements(w, spacer)

	w.writeEndElement()
}

func (sw *spacerWriter) writeAttributes(w *xmlWriter, spacer *models.Spacer) {
}

func (sw *spacerWriter) writeSubElements(w *xmlWriter, spacer *models.Spacer) {
}
