package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalKeyboard(e *etree.Element) (*models.Keyboard, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementKeyboard {
		return nil, &invalidTagError{
			expected: ElementKeyboard,
			actual:   e.Tag,
		}
	}

	kb := &models.Keyboard{}

	err := unmarshalKeyboardAttributes(kb, e.Attr)
	if err != nil {
		return nil, err
	}

	err = unmarshalKeyboardChildren(kb, e.Child)
	if err != nil {
		return nil, err
	}

	return kb, nil
}

func unmarshalKeyboardAttributes(kb *models.Keyboard, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeName:    {required: true},
		AttributeVersion: {required: false},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeName:
			kb.Name, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeVersion:
			kb.Version, err = unmarshalAttributeString(attr.Key, attr.Value)
		default:
			err = &unexpectedAttributeError{
				element:   ElementKeyboard,
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
				element:   ElementKeyboard,
				attribute: k,
			}
		}
	}

	return nil
}

func unmarshalKeyboardChildren(kb *models.Keyboard, children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementConstants:
			kb.Constants, err = unmarshalConstants(element)
		case ElementLayers:
			kb.Layers, err = unmarshalLayers(element)
		default:
			err = &invalidChildElementError{
				element: ElementKeyboard,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
