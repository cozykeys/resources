package models

var _ KeyboardElement = &Legend{}

// Legend TODO
type Legend struct {
	KeyboardElementBase

	Text                string
	FontSize            float64
	HorizontalAlignment LegendHorizontalAlignment
	VerticalAlignment   LegendVerticalAlignment
	Color               string
}

// LegendHorizontalAlignment TODO
type LegendHorizontalAlignment int

const (
	LegendHorizontalAlignmentLeft LegendHorizontalAlignment = iota
	LegendHorizontalAlignmentCenter
	LegendHorizontalAlignmentRight
)

var LegendHorizontalAlignmentStr = map[string]LegendHorizontalAlignment{
	"Left":   LegendHorizontalAlignmentLeft,
	"Center": LegendHorizontalAlignmentCenter,
	"Right":  LegendHorizontalAlignmentRight,
}

// LegendVerticalAlignment TODO
type LegendVerticalAlignment int

const (
	LegendVerticalAlignmentTop LegendVerticalAlignment = iota
	LegendVerticalAlignmentCenter
	LegendVerticalAlignmentBottom
)

var LegendVerticalAlignmentStr = map[string]LegendVerticalAlignment{
	"Top":    LegendVerticalAlignmentTop,
	"Center": LegendVerticalAlignmentCenter,
	"Bottom": LegendVerticalAlignmentBottom,
}
