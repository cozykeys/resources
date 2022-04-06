package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalAbsoluteMoveTo(e *etree.Element, parent models.KeyboardElement) (*models.AbsoluteMoveTo, error) {
	unmarshaller := &absoluteMoveToUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type absoluteMoveToUnmarshaller struct {
	element        *etree.Element
	absoluteMoveTo *models.AbsoluteMoveTo
	parent         models.KeyboardElement
}

func (u *absoluteMoveToUnmarshaller) unmarshal() (*models.AbsoluteMoveTo, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementAbsoluteMoveTo {
		return nil, &invalidTagError{
			expected: ElementAbsoluteMoveTo,
			actual:   u.element.Tag,
		}
	}

	u.absoluteMoveTo = &models.AbsoluteMoveTo{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.absoluteMoveTo.KeyboardElementBase); err != nil {
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

	return u.absoluteMoveTo, nil
}

func (u *absoluteMoveToUnmarshaller) unmarshalAttributes() error {
	return unmarshalElementAttributes(u.element, &u.absoluteMoveTo.KeyboardElementBase)
}

func (u *absoluteMoveToUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementEndPoint:
			u.absoluteMoveTo.EndPoint, err = unmarshalPoint(element, u.absoluteMoveTo)
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
