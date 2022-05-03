package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalCircle(e *etree.Element, parent models.KeyboardElement) (*models.Circle, error) {
	unmarshaller := &circleUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type circleUnmarshaller struct {
	element *etree.Element
	circle  *models.Circle
	parent  models.KeyboardElement
}

func (u *circleUnmarshaller) unmarshal() (*models.Circle, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementCircle {
		return nil, &invalidTagError{
			expected: ElementCircle,
			actual:   u.element.Tag,
		}
	}

	u.circle = &models.Circle{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.circle.KeyboardElementBase); err != nil {
		return nil, err
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	return u.circle, nil
}

func (u *circleUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeSize:        {required: true},
		AttributeFill:        {required: false},
		AttributeStroke:      {required: false},
		AttributeStrokeWidth: {required: false},
	}

	constants := u.circle.GetConstants()

	err := unmarshalElementAttributes(u.element, &u.circle.KeyboardElementBase)
	if err != nil {
		return err
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeSize:
			u.circle.Size, err = unmarshalAttributeFloat64(&attr, constants)
		case AttributeFill:
			u.circle.Fill, err = unmarshalAttributeString(&attr, constants)
		case AttributeStroke:
			u.circle.Stroke, err = unmarshalAttributeString(&attr, constants)
		case AttributeStrokeWidth:
			u.circle.StrokeWidth, err = unmarshalAttributeString(&attr, constants)
		default:
			if !isKeyboardElementAttribute(attr.Key) {
				err = &unexpectedAttributeError{
					element:   ElementLegend,
					attribute: attr.Key,
				}
			}
		}

		if err != nil {
			return err
		}

		if a, ok := supportedAttributes[attr.Key]; ok {
			a.found = true
		}
	}

	for k, v := range supportedAttributes {
		if v.required && v.found == false {
			return &missingRequiredAttributeError{
				element:   ElementLegend,
				attribute: k,
			}
		}
	}

	return nil
}
