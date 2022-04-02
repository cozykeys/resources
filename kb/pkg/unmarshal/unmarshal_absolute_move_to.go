package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalAbsoluteMoveTo(e *etree.Element) (*models.AbsoluteMoveTo, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementAbsoluteMoveTo {
		return nil, &invalidTagError{
			expected: ElementAbsoluteMoveTo,
			actual:   e.Tag,
		}
	}

	absoluteMoveTo := &models.AbsoluteMoveTo{}

	err := unmarshalAbsoluteMoveToChildren(absoluteMoveTo, e.Child)
	if err != nil {
		return nil, err
	}

	return absoluteMoveTo, nil
}

func unmarshalAbsoluteMoveToChildren(absoluteMoveTo *models.AbsoluteMoveTo, children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementEndPoint:
			absoluteMoveTo.EndPoint, err = unmarshalEndPoint(element)
		default:
			err = &invalidChildElementError{
				element: ElementAbsoluteMoveTo,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
