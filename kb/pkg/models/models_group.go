package models

// GroupChild TODO
type GroupChild interface{}

var _ KeyboardElement = &Group{}

// Group TODO
type Group struct {
	KeyboardElementBase

	Children []GroupChild
}
