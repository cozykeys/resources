package svg

import (
	"errors"
	"io"
	"kb/pkg/models"
)

func WriteLayer(w io.Writer, layer *models.Layer) error {
	writer := &layerWriter{
		writer: w,
	}

	return writer.Write()
}

type layerWriter struct {
	writer io.Writer

	/*
		g := parent.CreateElement("g")
		g.CreateAttr("id", key.Name)
		g.CreateAttr("transform", fmt.Sprintf("translate(%.3f,%.3f)", key.XOffset, key.YOffset))
	*/
	//public SvgGenerationOptions GenerationOptions { get; set; }
}

func (w *layerWriter) Write() error {
	return errors.New("not yet implemented")
}
