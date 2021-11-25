package kb

import (
	"log"
	"reflect"
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
