package models

var _ KeyboardElement = &AbsoluteVerticalLineTo{}
var _ PathComponent = &AbsoluteVerticalLineTo{}

// AbsoluteVerticalLineTo TODO
type AbsoluteVerticalLineTo struct {
	KeyboardElementBase

	Y float64
}

func (x *AbsoluteVerticalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
