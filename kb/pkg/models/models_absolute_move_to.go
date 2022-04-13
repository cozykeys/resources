package models

import "fmt"

var _ KeyboardElement = &AbsoluteMoveTo{}
var _ PathComponent = &AbsoluteMoveTo{}

// AbsoluteMoveTo TODO
type AbsoluteMoveTo struct {
	KeyboardElementBase

	EndPoint *Point
}

func (m *AbsoluteMoveTo) Data() string {
	return fmt.Sprintf("M %f %f", m.EndPoint.X, m.EndPoint.Y)
}
