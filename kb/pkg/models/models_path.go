package models

import "strings"

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

func (p *Path) Data() string {
	componentData := make([]string, len(p.Components))
	for i, component := range p.Components {
		componentData[i] = component.Data()
	}
	return strings.Join(componentData, " ")
}
