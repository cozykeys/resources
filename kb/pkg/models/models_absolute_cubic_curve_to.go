package models

var _ KeyboardElement = &AbsoluteCubicCurveTo{}
var _ PathComponent = &AbsoluteCubicCurveTo{}

// AbsoluteCubicCurveTo TODO
type AbsoluteCubicCurveTo struct {
	KeyboardElementBase

	EndPoint      Point
	ControlPointA Point
	ControlPointB Point
}

func (x *AbsoluteCubicCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
