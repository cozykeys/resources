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

	err = unmarshalLegendChildren(legend, e.Child)
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
		AttributeText:                {required: false},
		AttributeFontSize:            {required: false},
		AttributeColor:               {required: false},
		AttributeYOffset:             {required: false},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeHorizontalAlignment:
			legend.HorizontalAlignment, err = unmarshalLegendHorizontalAlignment(attr.Key, attr.Value)
		case AttributeVerticalAlignment:
			legend.VerticalAlignment, err = unmarshalLegendVerticalAlignment(attr.Key, attr.Value)
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

func unmarshalLegendChildren(legend *models.Legend, children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		// TODO
		default:
			err = &invalidChildElementError{
				element: ElementLegend,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
