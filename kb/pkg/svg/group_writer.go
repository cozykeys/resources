package svg

import (
	"kb/pkg/models"
	"log"

	"github.com/beevik/etree"
)

func writeGroup(parent *etree.Element, group *models.Group) error {
	writer := &groupWriter{
		parent: parent,
		group:  group,
	}

	return writer.write()
}

type groupWriter struct {
	parent *etree.Element
	group  *models.Group
}

func (w *groupWriter) write() error {
	if !w.group.Visible {
		return nil
	}

	g := w.parent.CreateElement("g")

	// Attributes
	g.CreateAttr("id", w.group.Name)
	writeTransform(g, w.group)

	// Child Elements
	if w.group.Debug {
		writeDebugOverlay(g, w.group)
	}

	for _, child := range w.group.Children {
		var err error
		switch v := child.(type) {
		case *models.Key:
			err = writeKey(g, v)
		case *models.Group:
			err = writeGroup(g, v)
		//case *models.Spacer:
		//	err = writeSpacer(g, v)
		//case *models.Stack:
		//	err = writeStack(g, v)
		case *models.Path:
			err = writePath(g, v)
		//case *models.Circle:
		//	err = writeCircle(g, v)
		//case *models.Text:
		//	err = writeText(g, v)
		default:
			log.Printf("Type not yet implemented %T", v)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
