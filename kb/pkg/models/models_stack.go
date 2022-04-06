package models

var _ KeyboardElement = &Stack{}

// Stack TODO
type Stack struct {
	Group

	Orientation StackOrientation
}

// TODO: Stack needs custom Width & Height methods

// StackOrientation TODO
type StackOrientation int

const (
	StackOrientationHorizontal StackOrientation = iota
	StackOrientationVertical
)

var StackOrientationStr = map[string]StackOrientation{
	"Horizontal": StackOrientationHorizontal,
	"Vertical":   StackOrientationVertical,
}
