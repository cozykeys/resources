package models

type PathComponent interface {
	Data() string
}

var _ KeyboardElement = &Path{}

// Path TODO
type Path struct {
	KeyboardElementBase

	Fill        string
	FillOpacity string
	Stroke      string
	StrokeWidth string
	Components  []PathComponent
}

func (x *Path) Data() string {
	//string Data => string.Join(" ", Components.Select(component => component.Data));
	return "NYI"
}
