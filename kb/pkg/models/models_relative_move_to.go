package models

var _ KeyboardElement = &RelativeMoveTo{}
var _ PathComponent = &RelativeMoveTo{}

// RelativeMoveTo TODO
type RelativeMoveTo struct {
	KeyboardElementBase

	EndPoint Point
}

func (x *RelativeMoveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}
