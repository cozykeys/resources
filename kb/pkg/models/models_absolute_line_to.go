package models

import "fmt"

var _ KeyboardElement = &AbsoluteLineTo{}
var _ PathComponent = &AbsoluteLineTo{}

// AbsoluteLineTo TODO
type AbsoluteLineTo struct {
	KeyboardElementBase

	EndPoint *Point
}

func (l *AbsoluteLineTo) Data() string {
	return fmt.Sprintf("L %f %f", l.EndPoint.X, l.EndPoint.Y)
}
