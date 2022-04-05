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
			Parent: u.parent,
		},
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
		AttributeXOffset:     {required: false},
		AttributeYOffset:     {required: false},
		AttributeFill:        {required: false},
		AttributeStroke:      {required: false},
		AttributeStrokeWidth: {required: false},
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeSize:
			u.circle.Size, err = unmarshalAttributeFloat64(&attr, u.circle.GetConstants())
		case AttributeXOffset:
			u.circle.XOffset, err = unmarshalAttributeFloat64(&attr, u.circle.GetConstants())
		case AttributeYOffset:
			u.circle.YOffset, err = unmarshalAttributeFloat64(&attr, u.circle.GetConstants())
		case AttributeFill:
			u.circle.Fill, err = unmarshalAttributeString(&attr, u.circle.GetConstants())
		case AttributeStroke:
			u.circle.Stroke, err = unmarshalAttributeString(&attr, u.circle.GetConstants())
		case AttributeStrokeWidth:
			u.circle.StrokeWidth, err = unmarshalAttributeString(&attr, u.circle.GetConstants())
		default:
			err = &unexpectedAttributeError{
				element:   ElementLegend,
				attribute: attr.Key,
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
