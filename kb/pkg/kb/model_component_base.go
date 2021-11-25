package kb

import "log"

type ComponentBase struct {
	Name    string
	XOffset float64
	YOffset float64
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
				return InvalidTypeError("ComponentBase", "Name")
			}
			cb.Name = valueStr
			requiredFields["Name"] = true
		case "XOffset":
			valueFloat, ok := value.(float64)
			if !ok {
				return InvalidTypeError("ComponentBase", "XOffset")
			}
			cb.XOffset = valueFloat
		case "YOffset":
			valueFloat, ok := value.(float64)
			if !ok {
				return InvalidTypeError("ComponentBase", "YOffset")
			}
			cb.YOffset = valueFloat
		}
	}

	for key, value := range requiredFields {
		if value == false {
			return MissingRequiredFieldError("ComponentBase", key)
		}
	}

	return nil
}
