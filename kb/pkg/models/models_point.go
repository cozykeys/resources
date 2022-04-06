package models

var _ KeyboardElement = &Point{}

// Point TODO
type Point struct {
	KeyboardElementBase

	X float64
	Y float64
}
