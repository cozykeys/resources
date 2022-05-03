package models

var _ KeyboardElement = &AbsoluteHorizontalLineTo{}
var _ PathComponent = &AbsoluteHorizontalLineTo{}

// AbsoluteHorizontalLineTo TODO
type AbsoluteHorizontalLineTo struct {
	KeyboardElementBase

	X float64
}

func (x *AbsoluteHorizontalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
