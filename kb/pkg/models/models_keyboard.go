package models

var _ KeyboardElement = &Keyboard{}

// Keyboard TODO
type Keyboard struct {
	KeyboardElementBase

	Version string
	Layers  []Layer
}
