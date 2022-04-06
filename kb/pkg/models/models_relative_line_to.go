package models

var _ KeyboardElement = &RelativeLineTo{}
var _ PathComponent = &RelativeLineTo{}

// RelativeLineTo TODO
type RelativeLineTo struct {
	KeyboardElementBase

	EndPoint Point
}

func (x *RelativeLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
