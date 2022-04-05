package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalLegend(e *etree.Element, parent models.KeyboardElement) (*models.Legend, error) {
	unmarshaller := &legendUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type legendUnmarshaller struct {
	element *etree.Element
	legend  *models.Legend
	parent  models.KeyboardElement
}

func (u *legendUnmarshaller) unmarshal() (*models.Legend, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementLegend {
		return nil, &invalidTagError{
			expected: ElementLegend,
			actual:   u.element.Tag,
		}
	}

	u.legend = &models.Legend{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent: u.parent,
		},
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	return u.legend, nil
}

func (u *legendUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeHorizontalAlignment: {required: false},
		AttributeVerticalAlignment:   {required: false},
		AttributeText:                {required: true},
		AttributeFontSize:            {required: false},
		AttributeColor:               {required: false},
		AttributeYOffset:             {required: false},
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeHorizontalAlignment:
			u.legend.HorizontalAlignment, err = unmarshalAttributeLegendHorizontalAlignment(&attr, u.legend.GetConstants())
		case AttributeVerticalAlignment:
			u.legend.VerticalAlignment, err = unmarshalAttributeLegendVerticalAlignment(&attr, u.legend.GetConstants())
		case AttributeText:
			u.legend.Text, err = unmarshalAttributeString(&attr, u.legend.GetConstants())
		case AttributeFontSize:
			u.legend.FontSize, err = unmarshalAttributeFloat64(&attr, u.legend.GetConstants())
		case AttributeColor:
			u.legend.Color, err = unmarshalAttributeString(&attr, u.legend.GetConstants())
		case AttributeYOffset:
			u.legend.YOffset, err = unmarshalAttributeFloat64(&attr, u.legend.GetConstants())
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
