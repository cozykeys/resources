package models

var _ KeyboardElement = &Text{}

// Text TODO
type Text struct {
	KeyboardElementBase

	Content    string
	TextAnchor string
	Font       string
	Fill       string
}
