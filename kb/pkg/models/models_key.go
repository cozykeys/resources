package models

var _ KeyboardElement = &Key{}

// Key TODO
type Key struct {
	KeyboardElementBase

	Legends []Legend
	Row     int
	Column  int
	Fill    string
	Stroke  string
}
