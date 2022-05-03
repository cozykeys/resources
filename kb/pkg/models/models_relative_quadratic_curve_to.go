package models

var _ KeyboardElement = &RelativeQuadraticCurveTo{}
var _ PathComponent = &RelativeQuadraticCurveTo{}

// RelativeQuadraticCurveTo TODO
type RelativeQuadraticCurveTo struct {
	KeyboardElementBase

	EndPoint     Point
	ControlPoint Point
}

func (x *RelativeQuadraticCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
