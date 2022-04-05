package svg

import (
	"errors"
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func WriteLayer(parent *etree.Element, layer *models.Layer) error {
	writer := &layerWriter{
		parent: parent,
	}

	return writer.Write()
}

type layerWriter struct {
	parent *etree.Element
}

func (w *layerWriter) Write() error {
	/*
		g := parent.CreateElement("g")
		g.CreateAttr("id", key.Name)
		g.CreateAttr("transform", fmt.Sprintf("translate(%.3f,%.3f)", key.XOffset, key.YOffset))
	*/
	//public SvgGenerationOptions GenerationOptions { get; set; }

	return errors.New("not yet implemented")
}
