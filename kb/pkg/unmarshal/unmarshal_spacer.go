package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalSpacer(e *etree.Element, parent models.KeyboardElement) (*models.Spacer, error) {
	unmarshaller := &spacerUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type spacerUnmarshaller struct {
	element *etree.Element
	spacer  *models.Spacer
	parent  models.KeyboardElement
}

func (u *spacerUnmarshaller) unmarshal() (*models.Spacer, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementSpacer {
		return nil, &invalidTagError{
			expected: ElementSpacer,
			actual:   u.element.Tag,
		}
	}

	u.spacer = &models.Spacer{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent: u.parent,
		},
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	return u.spacer, nil
}

func (u *spacerUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeHeight: {required: true},
		AttributeWidth:  {required: true},
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeHeight:
			u.spacer.Height, err = unmarshalAttributeFloat64(&attr, u.spacer.GetConstants())
		case AttributeWidth:
			u.spacer.Width, err = unmarshalAttributeFloat64(&attr, u.spacer.GetConstants())
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
