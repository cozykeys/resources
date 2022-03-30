package kb

import (
	"github.com/beevik/etree"
)

func ParseKeyboard(bytes []byte) (*Keyboard, error) {
	doc := etree.NewDocument()

	err := doc.ReadFromBytes(bytes)
	if err != nil {
		return nil, err
	}

	keyboard := &Keyboard{}
	err = keyboard.fromXMLElement(doc.Root())
	if err != nil {
		return nil, err
	}

	return keyboard, nil
}
