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
	/*
	   public void Write(XmlWriter writer, Group group)
	   {
	       elementWriter.WriteSubElements(writer, group);
	       WriteSubElements(writer, group);

	       writer.WriteEndElement();
	   }
	*/

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
		default:
			log.Printf("Type not yet implemented %T", v)
			/*
			   case var _ when child is Keyboard:
			       throw new InvalidDataException("Keyboard is not a valid child type.");
			   case Spacer spacer when child is Spacer:
			       var spacerWriter = new SpacerWriter { GenerationOptions = GenerationOptions };
			       spacerWriter.Write(writer, spacer);
			       break;
			   case Stack stack when child is Stack:
			       var stackWriter = new StackWriter { GenerationOptions = GenerationOptions };
			       stackWriter.Write(writer, stack);
			       break;
			   case Models.Path.Path path when child is Models.Path.Path:
			       var pathWriter = new PathWriter { GenerationOptions = GenerationOptions };
			       pathWriter.Write(writer, path);
			       break;
			   case Circle circle when child is Circle:
			       var holeWriter = new CircleWriter { GenerationOptions = GenerationOptions };
			       holeWriter.Write(writer, circle);
			       break;
			   case Text text when child is Text:
			       var textWriter = new TextWriter { GenerationOptions = GenerationOptions };
			       textWriter.Write(writer, text);
			       break;
			   case Group subGroup when child is Group:
			       Write(writer, subGroup);
			       break;
			   default:
			       throw new NotSupportedException();
			*/
		}

		if err != nil {
			return err
		}
	}

	return nil
}
