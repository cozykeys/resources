package svg

import (
	"fmt"
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func writeLegend(parent *etree.Element, legend *models.Legend) error {
	writer := &legendWriter{
		parent: parent,
		legend: legend,
	}

	return writer.write()
}

type legendWriter struct {
	parent *etree.Element
	legend *models.Legend
}

func (w *legendWriter) write() error {
	if !w.legend.Visible {
		return nil
	}

	g := w.parent.CreateElement("text")

	// Attributes
	g.CreateAttr("id", w.legend.Name)
	writeTransform(g, w.legend)

	// TODO: Is this necessary when the same thing is in the style?
	g.CreateAttr("text-anchor", "middle")

	fontSize := float64OrDefault(w.legend.FontSize, 4.0)

	g.CreateAttr("style", cssStyleString(map[string]string{
		"fill":              stringOrDefault(w.legend.Color, "#000000"),
		"dominant-baseline": "central",
		"text-anchor":       "middle",
		"font-size":         fmt.Sprintf("%fpx", fontSize),
		"font-family":       "sans-serif",
		"font-weight":       "normal",
		"font-style":        "normal",
	}))

	g.CreateText(w.legend.Text)
	return nil
}
