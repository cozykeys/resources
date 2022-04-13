package models

import "fmt"

var _ KeyboardElement = &AbsoluteQuadraticCurveTo{}
var _ PathComponent = &AbsoluteQuadraticCurveTo{}

// AbsoluteQuadraticCurveTo TODO
type AbsoluteQuadraticCurveTo struct {
	KeyboardElementBase

	EndPoint     Point
	ControlPoint Point
}

func (c *AbsoluteQuadraticCurveTo) Data() string {
	return fmt.Sprintf("Q %f %f %f %f", c.ControlPoint.X, c.ControlPoint.Y, c.EndPoint.X, c.EndPoint.Y)
}
