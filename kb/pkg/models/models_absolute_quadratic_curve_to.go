package models

var _ KeyboardElement = &AbsoluteQuadraticCurveTo{}
var _ PathComponent = &AbsoluteQuadraticCurveTo{}

// AbsoluteQuadraticCurveTo TODO
type AbsoluteQuadraticCurveTo struct {
	KeyboardElementBase

	EndPoint     Point
	ControlPoint Point
}

func (x *AbsoluteQuadraticCurveTo) Data() string {
	//string Data => $"Q {ControlPoint.X} {ControlPoint.Y} {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}
