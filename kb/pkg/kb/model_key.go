package kb

import (
	"log"
	"reflect"
)

type Key struct {
	ComponentBase

	Legends []Legend
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
				return InvalidTypeError("Key", "Legends")
			}

			// TODO: Would be nice to have a "fromSlice" or something similar
			// to "fromMap" to reduce LoC in this func

			for _, legend := range legends {
				legendMap, ok := legend.(map[string]interface{})
				if !ok {
					return InvalidTypeError("Key", "Legends")
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
			return MissingRequiredFieldError("Key", k)
		}
	}

	log.Printf("(*Key)::fromMap()")
	return nil
}
