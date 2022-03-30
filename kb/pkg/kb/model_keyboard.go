package kb

import (
	"log"
	"reflect"
	"strconv"

	"github.com/beevik/etree"
)

const (
	ElementKeyboard = "Keyboard"

	AttributeName    = "Name"
	AttributeVersion = "Version"
	AttributeWidth   = "Width"
	AttributeHeight  = "Height"
)

type Keyboard struct {
	Name       string
	Version    string
	Width      float64
	Height     float64
	Components []Component
}

func (kb *Keyboard) fromMap(m map[string]interface{}) error {
	log.Printf("(*Keyboard)::fromMap()")

	requiredFields := map[string]bool{
		"Name":       false,
		"Version":    false,
		"Height":     false,
		"Width":      false,
		"Components": false,
	}

	for key, value := range m {
		switch key {
		case "Name":
			valueStr, ok := value.(string)
			if !ok {
				return InvalidTypeError("Keyboard", "Name")
			}
			kb.Name = valueStr
			requiredFields["Name"] = true
		case "Version":
			valueStr, ok := value.(string)
			if !ok {
				return InvalidTypeError("Keyboard", "Version")
			}
			kb.Version = valueStr
			requiredFields["Version"] = true
		case "Height":
			valueFloat, ok := value.(float64)
			if !ok {
				return InvalidTypeError("Keyboard", "Height")
			}
			kb.Height = valueFloat
			requiredFields["Height"] = true
		case "Width":
			valueFloat, ok := value.(float64)
			if !ok {
				return InvalidTypeError("Keyboard", "Width")
			}
			kb.Width = valueFloat
			requiredFields["Width"] = true
		case "Components":
			cmps, ok := value.([]interface{})
			if !ok {
				log.Printf("%s", reflect.TypeOf(value))
				return InvalidTypeError("Keyboard", "Components")
			}

			// TODO: Would be nice to have a "fromSlice" or something similar
			// to "fromMap" to reduce LoC in this func

			for _, cmp := range cmps {
				cmpMap, ok := cmp.(map[string]interface{})
				if !ok {
					return InvalidTypeError("Keyboard", "Components")
				}

				t, ok := cmpMap["Type"]
				if !ok {
					return MissingRequiredFieldError("Component", "Type")
				}

				switch t {
				case "Key":
					key := &Key{}
					if err := key.fromMap(cmpMap); err != nil {
						return err
					}
					kb.Components = append(kb.Components, key)
				case "Circle":
					circle := &Circle{}
					if err := circle.fromMap(cmpMap); err != nil {
						return err
					}
					kb.Components = append(kb.Components, circle)
				}
			}

			requiredFields["Components"] = true
		default:
			return UnexpectedFieldError("Keyboard", key)
		}
	}

	for key, value := range requiredFields {
		if value == false {
			return MissingRequiredFieldError("Keyboard", key)
		}
	}

	// TODO: Do any post-processing such as ensuring required fields were set
	// and values are valid

	return nil
}

func (kb *Keyboard) fromXMLElement(e *etree.Element) error {
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

		switch attr.Key {
		case AttributeName:
			kb.Name = attr.Value
		case AttributeVersion:
			kb.Version = attr.Value
		case AttributeHeight:
			height, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return &invalidAttributeTypeError{
					element:   ElementKeyboard,
					attribute: attr.Key,
				}
			}
			kb.Height = height
		case AttributeWidth:
			width, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return &invalidAttributeTypeError{
					element:   ElementKeyboard,
					attribute: attr.Key,
				}
			}
			kb.Height = width
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
