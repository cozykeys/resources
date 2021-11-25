package kb

import "log"

type Circle struct {
	ComponentBase

	Radius float64
}

func (circle *Circle) fromMap(m map[string]interface{}) error {
	if err := circle.ComponentBase.fromMap(m); err != nil {
		return err
	}

	log.Printf("(*Circle)::fromMap()")
	return nil
}
