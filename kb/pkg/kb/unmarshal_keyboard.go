package kb

import (
	"log"

	"github.com/beevik/etree"
)

const (
	ElementKeyboard = "Keyboard"

	AttributeName    = "Name"
	AttributeVersion = "Version"
	AttributeWidth   = "Width"
	AttributeHeight  = "Height"
)

func (kb *Keyboard) unmarshal(e *etree.Element) error {
	if e == nil {
		return &nilElementError{}
	}

	if e.Tag != ElementKeyboard {
		return &invalidTagError{
			expected: ElementKeyboard,
			actual:   e.Tag,
		}
	}

	err := kb.parseAttributes(e.Attr)
	if err != nil {
		return err
	}

	err = kb.parseChildren(e.Child)
	if err != nil {
		return err
	}

	return nil
}

func (kb *Keyboard) parseAttributes(attributes []etree.Attr) error {
	supportedAttributes := map[string]struct {
		required bool
		found    bool
	}{
		AttributeName:    {required: true},
		AttributeVersion: {required: false},
		AttributeWidth:   {required: false},
		AttributeHeight:  {required: false},
	}

	for _, attr := range attributes {
		log.Printf("Attribute key = %q, value = %q\n", attr.Key, attr.Value)

		var err error
		switch attr.Key {
		case AttributeName:
			kb.Name, err = unmarshalAttributeString(attr.Key, attr.Value)
			if err != nil {
				return err
			}
		case AttributeVersion:
			kb.Version, err = unmarshalAttributeString(attr.Key, attr.Value)
			if err != nil {
				return err
			}
		case AttributeHeight:
			kb.Height, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
			if err != nil {
				return err
			}
		case AttributeWidth:
			kb.Width, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
			if err != nil {
				return err
			}
		default:
			return &unexpectedAttributeError{
				element:   ElementKeyboard,
				attribute: attr.Key,
			}
		}

		if a, ok := supportedAttributes[attr.Key]; ok && a.required {
			a.found = true
		}
	}

	for k, v := range supportedAttributes {
		if v.found == false {
			return &missingRequiredAttributeError{
				element:   ElementKeyboard,
				attribute: k,
			}
		}
	}

	return nil
}

func (kb *Keyboard) parseChildren(children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			// Ignore children that are not elements
			continue
		}

		switch element.Tag {
		case "Foo":
			log.Print("Foo")
		default:
			return &invalidChildElementError{
				element: ElementKeyboard,
				child:   element.Tag,
			}
		}
	}

	return nil
}
