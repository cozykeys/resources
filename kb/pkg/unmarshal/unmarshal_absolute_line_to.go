package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalAbsoluteLineTo(e *etree.Element, parent models.KeyboardElement) (*models.AbsoluteLineTo, error) {
	unmarshaller := &absoluteLineToUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type absoluteLineToUnmarshaller struct {
	element        *etree.Element
	absoluteLineTo *models.AbsoluteLineTo
	parent         models.KeyboardElement
}

func (u *absoluteLineToUnmarshaller) unmarshal() (*models.AbsoluteLineTo, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementAbsoluteLineTo {
		return nil, &invalidTagError{
			expected: ElementAbsoluteLineTo,
			actual:   u.element.Tag,
		}
	}

	u.absoluteLineTo = &models.AbsoluteLineTo{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.absoluteLineTo.KeyboardElementBase); err != nil {
		return nil, err
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	err = u.unmarshalChildElements()
	if err != nil {
		return nil, err
	}

	return u.absoluteLineTo, nil
}

func (u *absoluteLineToUnmarshaller) unmarshalAttributes() error {
	return unmarshalElementAttributes(u.element, &u.absoluteLineTo.KeyboardElementBase)
}

func (u *absoluteLineToUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementEndPoint:
			u.absoluteLineTo.EndPoint, err = unmarshalPoint(element, u.absoluteLineTo)
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
