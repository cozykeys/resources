package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalText(e *etree.Element, parent models.KeyboardElement) (*models.Text, error) {
	unmarshaller := &textUnmarshaller{
		element: e,
		parent:  parent,
	}

	return unmarshaller.unmarshal()
}

type textUnmarshaller struct {
	element *etree.Element
	text    *models.Text
	parent  models.KeyboardElement
}

func (u *textUnmarshaller) unmarshal() (*models.Text, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementText {
		return nil, &invalidTagError{
			expected: ElementText,
			actual:   u.element.Tag,
		}
	}

	u.text = &models.Text{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.text.KeyboardElementBase); err != nil {
		return nil, err
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	return u.text, nil
}

func (u *textUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeContent:    {required: false},
		AttributeTextAnchor: {required: false},
		AttributeFont:       {required: false},
		AttributeFill:       {required: false},
	}

	constants := u.text.GetConstants()

	err := unmarshalElementAttributes(u.element, &u.text.KeyboardElementBase)
	if err != nil {
		return err
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeContent:
			u.text.Content, err = unmarshalAttributeString(&attr, constants)
		case AttributeTextAnchor:
			u.text.TextAnchor, err = unmarshalAttributeString(&attr, constants)
		case AttributeFont:
			u.text.Font, err = unmarshalAttributeString(&attr, constants)
		case AttributeFill:
			u.text.Fill, err = unmarshalAttributeString(&attr, constants)
		default:
			if !isKeyboardElementAttribute(attr.Key) {
				err = &unexpectedAttributeError{
					element:   ElementText,
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
				element:   ElementText,
				attribute: k,
			}
		}
	}

	return nil
}
