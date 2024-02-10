package models

// GroupChild TODO
type GroupChild interface {
	KeyboardElement
}

var _ KeyboardElement = &Group{}

// Group TODO
type Group struct {
	KeyboardElementBase

	Children []GroupChild
}

type IGroup interface {
	GetElement() KeyboardElement
	GetChildren() []GroupChild
}

func (g *Group) GetElement() KeyboardElement {
	return g
}

func (g *Group) GetChildren() []GroupChild {
	return g.Children
}
