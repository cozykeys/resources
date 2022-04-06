package models

var _ KeyboardElement = &Constant{}

// Constant TODO
type Constant struct {
	KeyboardElementBase

	Value string
}
