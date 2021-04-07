package kbutil

import (
	"encoding/json"
	"fmt"
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

func (key *Key) fromMap(m map[string]interface{}) error {
	log.Printf("(*Key)::fromMap()")
	return nil
}

func (circle *Circle) fromMap(m map[string]interface{}) error {
	log.Printf("(*Circle)::fromMap()")
	return nil
}

func (kb *Keyboard) fromMap(m map[string]interface{}) error {
	log.Printf("(*Keyboard)::fromMap()")

	requiredFields := map[string]bool{
		"Name":       false,
		"Version":    false,
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

func (k *Key) UnmarshalJSON(bytes []byte) error {
	fmt.Printf("(*Key)::UnmarshalJSON()\n")
	return nil
}

func (c *Circle) UnmarshalJSON(bytes []byte) error {
	fmt.Printf("(*Circle)::UnmarshalJSON()\n")
	return nil
}
