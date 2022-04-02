package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalAbsoluteLineTo(e *etree.Element) (*models.AbsoluteLineTo, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementAbsoluteLineTo {
		return nil, &invalidTagError{
			expected: ElementAbsoluteLineTo,
			actual:   e.Tag,
		}
	}

	absoluteLineTo := &models.AbsoluteLineTo{}

	err := unmarshalAbsoluteLineToChildren(absoluteLineTo, e.Child)
	if err != nil {
		return nil, err
	}

	return absoluteLineTo, nil
}

func unmarshalAbsoluteLineToChildren(absoluteLineTo *models.AbsoluteLineTo, children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementEndPoint:
			absoluteLineTo.EndPoint, err = unmarshalEndPoint(element)
		default:
			err = &invalidChildElementError{
				element: ElementAbsoluteLineTo,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
