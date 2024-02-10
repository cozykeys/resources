package svg

import (
	"kb/pkg/models"
)

type layerWriter struct {
	options *Options
}

func (lw *layerWriter) write(w *xmlWriter, layer *models.Layer) {
	if !layer.Visible {
		return
	}

	w.writeStartElement("g")

	ew := &elementWriter{options: lw.options}

	ew.writeAttributes(w, layer)
	lw.writeAttributes(w, layer)

	ew.writeSubElements(w, layer)
	lw.writeSubElements(w, layer)

	w.writeEndElement()
}

func (lw *layerWriter) writeAttributes(w *xmlWriter, layer *models.Layer) {
}

func (lw *layerWriter) writeSubElements(w *xmlWriter, layer *models.Layer) {
	lw.writeGroups(w, layer)
}

func (lw *layerWriter) writeGroups(w *xmlWriter, layer *models.Layer) {
	gw := &groupWriter{options: lw.options}
	for _, group := range layer.Groups {
		gw.write(w, &group)
	}
}
