package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalKeyboard(e *etree.Element) (*models.Keyboard, error) {
	unmarshaller := &keyboardUnmarshaller{element: e}
	return unmarshaller.unmarshal()
}

type keyboardUnmarshaller struct {
	element  *etree.Element
	keyboard *models.Keyboard
}

func (u *keyboardUnmarshaller) unmarshal() (*models.Keyboard, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementKeyboard {
		return nil, &invalidTagError{
			expected: ElementKeyboard,
			actual:   u.element.Tag,
		}
	}

	u.keyboard = &models.Keyboard{}

	if err := u.unmarshalConstants(); err != nil {
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

	return u.keyboard, nil
}

func (u *keyboardUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeName:    {required: true},
		AttributeVersion: {required: false},
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeName:
			u.keyboard.Name, err = unmarshalAttributeString(&attr, u.keyboard.GetConstants())
		case AttributeVersion:
			u.keyboard.Version, err = unmarshalAttributeString(&attr, u.keyboard.GetConstants())
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

func (u *keyboardUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementConstants:
			// Skip this here; it was already processed
		case ElementLayers:
			u.keyboard.Layers, err = unmarshalLayers(element, u.keyboard)
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

func (u *keyboardUnmarshaller) unmarshalConstants() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		if element.Tag == ElementConstants {
			var err error
			u.keyboard.Constants, err = unmarshalConstants(element, u.keyboard)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}
