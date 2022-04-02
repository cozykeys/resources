package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalKey(e *etree.Element) (*models.Key, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementKey {
		return nil, &invalidTagError{
			expected: ElementKey,
			actual:   e.Tag,
		}
	}

	key := &models.Key{}

	err := unmarshalKeyAttributes(key, e.Attr)
	if err != nil {
		return nil, err
	}

	err = unmarshalKeyChildren(key, e.Child)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func unmarshalKeyAttributes(key *models.Key, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeName:    {required: false},
		AttributeRow:     {required: false},
		AttributeColumn:  {required: false},
		AttributeXOffset: {required: false},
		AttributeYOffset: {required: false},
		AttributeWidth:   {required: false},
		AttributeHeight:  {required: false},
		AttributeMargin:  {required: false},
		AttributeFill:    {required: false},
		AttributeStroke:  {required: false},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeName:
			key.Name, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeRow:
			key.Row, err = unmarshalAttributeInt(attr.Key, attr.Value)
		case AttributeColumn:
			key.Column, err = unmarshalAttributeInt(attr.Key, attr.Value)
		case AttributeXOffset:
			key.XOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeYOffset:
			key.YOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeWidth:
			key.Width, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeHeight:
			key.Height, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeMargin:
			key.Margin, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeFill:
			key.Fill, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeStroke:
			key.Stroke, err = unmarshalAttributeString(attr.Key, attr.Value)
		default:
			err = &unexpectedAttributeError{
				element:   ElementKey,
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
				element:   ElementKey,
				attribute: k,
			}
		}
	}

	return nil
}

func unmarshalKeyChildren(key *models.Key, children []etree.Token) error {
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
				element: ElementKey,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
