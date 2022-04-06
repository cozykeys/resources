package models

var _ KeyboardElement = &AbsoluteMoveTo{}
var _ PathComponent = &AbsoluteMoveTo{}

// AbsoluteMoveTo TODO
type AbsoluteMoveTo struct {
	KeyboardElementBase

	EndPoint *Point
}

func (x *AbsoluteMoveTo) Data() string {
	//string Data => $"M {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}
