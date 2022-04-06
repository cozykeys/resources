package models

var _ KeyboardElement = &AbsoluteLineTo{}
var _ PathComponent = &AbsoluteLineTo{}

// AbsoluteLineTo TODO
type AbsoluteLineTo struct {
	KeyboardElementBase

	EndPoint *Point
}

func (x *AbsoluteLineTo) Data() string {
	//string Data => $"L {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}
