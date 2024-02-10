package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalAbsoluteQuadraticCurveTo(e *etree.Element, parent models.KeyboardElement) (*models.AbsoluteQuadraticCurveTo, error) {
	unmarshaller := &absoluteQuadraticCurveToUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type absoluteQuadraticCurveToUnmarshaller struct {
	element                  *etree.Element
	absoluteQuadraticCurveTo *models.AbsoluteQuadraticCurveTo
	parent                   models.KeyboardElement
}

func (u *absoluteQuadraticCurveToUnmarshaller) unmarshal() (*models.AbsoluteQuadraticCurveTo, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementAbsoluteQuadraticCurveTo {
		return nil, &invalidTagError{
			expected: ElementAbsoluteQuadraticCurveTo,
			actual:   u.element.Tag,
		}
	}

	u.absoluteQuadraticCurveTo = &models.AbsoluteQuadraticCurveTo{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.absoluteQuadraticCurveTo.KeyboardElementBase); err != nil {
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

	return u.absoluteQuadraticCurveTo, nil
}

func (u *absoluteQuadraticCurveToUnmarshaller) unmarshalAttributes() error {
	return unmarshalElementAttributes(u.element, &u.absoluteQuadraticCurveTo.KeyboardElementBase)
}

func (u *absoluteQuadraticCurveToUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementEndPoint:
			u.absoluteQuadraticCurveTo.EndPoint, err = unmarshalPoint(element, u.absoluteQuadraticCurveTo)
		case ElementControlPoint:
			u.absoluteQuadraticCurveTo.ControlPoint, err = unmarshalPoint(element, u.absoluteQuadraticCurveTo)
		default:
			err = &invalidChildElementError{
				element: u.element.Tag,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
