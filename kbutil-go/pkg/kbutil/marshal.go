package kbutil

import (
	"encoding/json"
	"log"
)

func (kb *Keyboard) MarshalJSON() ([]byte, error) {
	log.Printf("(*Keyboard)::MarshalJSON()")

	m, err := kb.toMap()
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (kb *Keyboard) toMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"Name":       kb.Name,
		"Version":    kb.Version,
		"Components": "TODO",
	}

	return m, nil
}
