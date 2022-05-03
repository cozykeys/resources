package svg

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func writeLayer(parent *etree.Element, layer *models.Layer) error {
	writer := &layerWriter{
		parent: parent,
		layer:  layer,
	}

	return writer.write()
}

type layerWriter struct {
	parent *etree.Element
	layer  *models.Layer
}

func (w *layerWriter) write() error {
	if !w.layer.Visible {
		return nil
	}

	g := w.parent.CreateElement("g")

	// Attributes
	g.CreateAttr("id", w.layer.Name)
	writeTransform(g, w.layer)

	// Child Elements
	if w.layer.Debug {
		writeDebugOverlay(g, w.layer)
	}

	for _, group := range w.layer.Groups {
		err := writeGroup(g, &group)
		if err != nil {
			return err
		}
	}

	return nil
}
