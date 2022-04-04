package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

/*
   "Text": {
       "Attributes": [
           "Content",
           "TextAnchor",
           "Font",
           "Fill",
           "XOffset",
           "YOffset"
       ],
       "Children": []
   }
*/

func unmarshalText(e *etree.Element) (*models.Text, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementText {
		return nil, &invalidTagError{
			expected: ElementText,
			actual:   e.Tag,
		}
	}

	text := &models.Text{}

	err := unmarshalTextAttributes(text, e.Attr)
	if err != nil {
		return nil, err
	}

	return text, nil
}

func unmarshalTextAttributes(text *models.Text, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeContent:    {required: false},
		AttributeTextAnchor: {required: false},
		AttributeFont:       {required: false},
		AttributeFill:       {required: false},
		AttributeXOffset:    {required: false},
		AttributeYOffset:    {required: false},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeContent:
			text.Content, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeTextAnchor:
			text.TextAnchor, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeFont:
			text.Font, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeFill:
			text.Fill, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeXOffset:
			text.XOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeYOffset:
			text.YOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		default:
			err = &unexpectedAttributeError{
				element:   ElementText,
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
				element:   ElementText,
				attribute: k,
			}
		}
	}

	return nil
}
