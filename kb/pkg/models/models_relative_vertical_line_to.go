package models

var _ KeyboardElement = &RelativeVerticalLineTo{}
var _ PathComponent = &RelativeVerticalLineTo{}

// RelativeVerticalLineTo TODO
type RelativeVerticalLineTo struct {
	KeyboardElementBase

	Y float64
}

func (x *RelativeVerticalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
