package models

var _ KeyboardElement = &Layer{}

// Layer TODO
type Layer struct {
	KeyboardElementBase

	ZIndex int
	Groups []Group
}
