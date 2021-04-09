package kbutil

import (
	"encoding/json"
	"log"
	"reflect"
)

// TODO: Explanation; this is the only exported method related to unmarshaling.
// User of the package is expected to unmarshal a la:
//
//     kb := &kbutil.Keyboard{}
//     if err := json.Unmarshal(bytes, kb); err != nil { ... }
//
// Internally, all unmarshaling is done by handle from generic maps/slices. The
// primary reason for this is that the schema is somewhat complicated and
// requires overriding some of the unmarshal functions anyways to support:
// - Sum types
// - Required fields
// As an added benefit, we don't have to unmarshal anything twice so it ends up
// being a bit more performant, not that that matters much.
func (kb *Keyboard) UnmarshalJSON(bytes []byte) error {
	kbMap := map[string]interface{}{}
	if err := json.Unmarshal(bytes, &kbMap); err != nil {
		return err
	}

	if err := kb.fromMap(kbMap); err != nil {
		return err
	}

	return nil
}

func (cb *ComponentBase) fromMap(m map[string]interface{}) error {
	log.Printf("(*ComponentBase)::fromMap()")

	requiredFields := map[string]bool{
		"Name": false,
	}

	for key, value := range m {
		switch key {
		case "Name":
			valueStr, ok := value.(string)
			if !ok {
				return errInvalidType("ComponentBase", "Name")
			}
			cb.Name = valueStr
			requiredFields["Name"] = true
		case "XOffset":
			valueFloat, ok := value.(float64)
			if !ok {
				return errInvalidType("ComponentBase", "XOffset")
			}
			cb.XOffset = valueFloat
		case "YOffset":
			valueFloat, ok := value.(float64)
			if !ok {
				return errInvalidType("ComponentBase", "YOffset")
			}
			cb.YOffset = valueFloat
		}
	}

	for key, value := range requiredFields {
		if value == false {
			return errMissingRequired("ComponentBase", key)
		}
	}

	return nil
}

func (legend *Legend) fromMap(m map[string]interface{}) error {
	return nil
}

func (key *Key) fromMap(m map[string]interface{}) error {
	if err := key.ComponentBase.fromMap(m); err != nil {
		return err
	}

	requiredFields := map[string]bool{
		"Legends": false,
	}

	for k, v := range m {
		switch k {
		case "Legends":
			legends, ok := v.([]interface{})
			if !ok {
				log.Printf("%s", reflect.TypeOf(v))
				return errInvalidType("Key", "Legends")
			}

			// TODO: Would be nice to have a "fromSlice" or something similar
			// to "fromMap" to reduce LoC in this func

			for _, legend := range legends {
				legendMap, ok := legend.(map[string]interface{})
				if !ok {
					return errInvalidType("Key", "Legends")
				}

				l := Legend{}
				if err := l.fromMap(legendMap); err != nil {
					return err
				}

				key.Legends = append(key.Legends, l)
			}

			requiredFields["Legends"] = true
		}
	}

	for k, v := range requiredFields {
		if v == false {
			return errMissingRequired("Key", k)
		}
	}

	log.Printf("(*Key)::fromMap()")
	return nil
}

func (circle *Circle) fromMap(m map[string]interface{}) error {
	if err := circle.ComponentBase.fromMap(m); err != nil {
		return err
	}

	log.Printf("(*Circle)::fromMap()")
	return nil
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
				return errInvalidType("Keyboard", "Name")
			}
			kb.Name = valueStr
			requiredFields["Name"] = true
		case "Version":
			valueStr, ok := value.(string)
			if !ok {
				return errInvalidType("Keyboard", "Version")
			}
			kb.Version = valueStr
			requiredFields["Version"] = true
		case "Height":
			valueFloat, ok := value.(float64)
			if !ok {
				return errInvalidType("Keyboard", "Height")
			}
			kb.Height = valueFloat
			requiredFields["Height"] = true
		case "Width":
			valueFloat, ok := value.(float64)
			if !ok {
				return errInvalidType("Keyboard", "Width")
			}
			kb.Width = valueFloat
			requiredFields["Width"] = true
		case "Components":
			cmps, ok := value.([]interface{})
			if !ok {
				log.Printf("%s", reflect.TypeOf(value))
				return errInvalidType("Keyboard", "Components")
			}

			// TODO: Would be nice to have a "fromSlice" or something similar
			// to "fromMap" to reduce LoC in this func

			for _, cmp := range cmps {
				cmpMap, ok := cmp.(map[string]interface{})
				if !ok {
					return errInvalidType("Keyboard", "Components")
				}

				t, ok := cmpMap["Type"]
				if !ok {
					return errMissingRequired("Component", "Type")
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
			return errUnexpectedField("Keyboard", key)
		}
	}

	for key, value := range requiredFields {
		if value == false {
			return errMissingRequired("Keyboard", key)
		}
	}

	// TODO: Do any post-processing such as ensuring required fields were set
	// and values are valid

	return nil
}
