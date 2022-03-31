package kb

import (
	"strconv"

	"github.com/beevik/etree"
)

func Unmarshal(bytes []byte) (*Keyboard, error) {
	doc := etree.NewDocument()

	err := doc.ReadFromBytes(bytes)
	if err != nil {
		return nil, err
	}

	keyboard := &Keyboard{}
	err = keyboard.unmarshal(doc.Root())
	if err != nil {
		return nil, err
	}

	return keyboard, nil
}

func unmarshalAttributeString(key, raw string) (string, error) {
	// TODO: Process constants
	return raw, nil
}

func unmarshalAttributeFloat64(key, raw string) (float64, error) {
	// TODO: Process constants
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, &invalidAttributeTypeError{
			element:   ElementKeyboard,
			attribute: key,
		}
	}
	return val, nil
}
