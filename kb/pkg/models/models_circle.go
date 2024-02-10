package models

var _ KeyboardElement = &Circle{}

// Circle TODO
type Circle struct {
	KeyboardElementBase

	Size        float64
	Fill        string
	Stroke      string
	StrokeWidth string
}

func (p *Circle) Data() string {
	return "TODO"
}
