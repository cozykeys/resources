package svg

import (
	"kb/pkg/models"
	"log"

	"github.com/beevik/etree"
)

const (
	DefaultPathFill        = "none"
	DefaultPathFillOpacity = "1"
	DefaultPathStroke      = "#000000"
	DefaultPathStrokeWidth = "0.5"
)

func writePath(parent *etree.Element, path *models.Path) error {
	writer := &pathWriter{
		parent: parent,
		path:   path,
	}

	return writer.write()
}

type pathWriter struct {
	parent *etree.Element
	path   *models.Path
}

func (w *pathWriter) write() error {
	if !w.path.Visible {
		return nil
	}

	g := w.parent.CreateElement("g")

	// Attributes
	g.CreateAttr("id", w.path.Name)
	writeTransform(g, w.path)

	// Elements
	if w.path.Debug {
		log.Printf("Writing debug overlay...")
		writeDebugOverlay(g, w.path)
	}

	path := w.parent.CreateElement("path")
	path.CreateAttr("style", cssStyleString(map[string]string{
		"fill":         stringOrDefault(w.path.Fill, "none"),
		"fill-opacity": stringOrDefault(w.path.FillOpacity, "1"),
		"stroke":       stringOrDefault(w.path.Stroke, "#000000"),
		"stroke-width": stringOrDefault(w.path.StrokeWidth, "0.5"),
	}))
	path.CreateAttr("d", w.path.Data())

	return nil
}
