package models

var _ KeyboardElement = &RelativeHorizontalLineTo{}
var _ PathComponent = &RelativeHorizontalLineTo{}

// RelativeHorizontalLineTo TODO
type RelativeHorizontalLineTo struct {
	KeyboardElementBase

	X float64
}

func (x *RelativeHorizontalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
