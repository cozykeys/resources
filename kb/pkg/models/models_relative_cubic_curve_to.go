package models

var _ KeyboardElement = &RelativeCubicCurveTo{}
var _ PathComponent = &RelativeCubicCurveTo{}

// RelativeCubicCurveTo TODO
type RelativeCubicCurveTo struct {
	KeyboardElementBase

	EndPoint      Point
	ControlPointA Point
	ControlPointB Point
}

func (x *RelativeCubicCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
