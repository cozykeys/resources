package svg

import (
	"kb/pkg/models"
	"log"
)

type groupWriter struct {
	options *Options
}

func newGroupWriter(options *Options) *groupWriter {
	return &groupWriter{options: options}
}

func (gw *groupWriter) write(w *xmlWriter, group models.IGroup) {
	if !group.GetElement().GetVisible() {
		return
	}

	w.writeStartElement("g")

	_elementWriter := newElementWriter(gw.options)

	_elementWriter.writeAttributes(w, group.GetElement())
	gw.writeAttributes(w, group)

	_elementWriter.writeSubElements(w, group.GetElement())
	gw.writeSubElements(w, group)

	w.writeEndElement()
}

func (gw *groupWriter) writeAttributes(w *xmlWriter, group models.IGroup) {
}

func (gw *groupWriter) writeSubElements(w *xmlWriter, group models.IGroup) {
	for _, child := range group.GetChildren() {
		switch v := child.(type) {
		case *models.Circle:
			_circleWriter := newCircleWriter(gw.options)
			_circleWriter.write(w, v)
		case *models.Group:
			_groupWriter := newGroupWriter(gw.options)
			_groupWriter.write(w, v)
		case *models.Key:
			_keyWriter := newKeyWriter(gw.options)
			_keyWriter.write(w, v)
		case *models.Path:
			_pathWriter := newPathWriter(gw.options)
			_pathWriter.write(w, v)
		case *models.Spacer:
			_spacerWriter := newSpacerWriter(gw.options)
			_spacerWriter.write(w, v)
		case *models.Stack:
			_stackWriter := newStackWriter(gw.options)
			_stackWriter.write(w, v)
		case *models.Text:
			_textWriter := newTextWriter(gw.options)
			_textWriter.write(w, v)
		default:
			log.Printf("Type not yet implemented %T", v)
		}
	}
}
