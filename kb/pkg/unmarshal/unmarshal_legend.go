package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalLegend(e *etree.Element) (*models.Legend, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementLegend {
		return nil, &invalidTagError{
			expected: ElementLegend,
			actual:   e.Tag,
		}
	}

	legend := &models.Legend{}

	err := unmarshalLegendAttributes(legend, e.Attr)
	if err != nil {
		return nil, err
	}

	return legend, nil
}

func unmarshalLegendAttributes(legend *models.Legend, attributes []etree.Attr) error {
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

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeHorizontalAlignment:
			legend.HorizontalAlignment, err = unmarshalAttributeLegendHorizontalAlignment(&attr)
		case AttributeVerticalAlignment:
			legend.VerticalAlignment, err = unmarshalAttributeLegendVerticalAlignment(&attr)
		case AttributeText:
			legend.Text, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeFontSize:
			legend.FontSize, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeColor:
			legend.Color, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeYOffset:
			legend.YOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
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
