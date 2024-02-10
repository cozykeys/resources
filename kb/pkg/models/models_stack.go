package models

import (
	"log"
	"math"
)

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

func (s *Stack) GetWidth() float64 {
	log.Print("Stack::GetStackWidth")
	if s.Orientation == StackOrientationHorizontal {
		width := 0.0
		for _, child := range s.Children {
			width += getTotalWidth(child)
		}
		return width
	} else if s.Orientation == StackOrientationVertical {
		minX := math.MaxFloat64
		maxX := -math.MaxFloat64
		for _, child := range s.Children {
			minX = F64Min(minX, getMinX(child))
			maxX = F64Max(maxX, getMaxX(child))
		}
		return maxX - minX
	}

	panic("unknown stack orientation")
}

func (s *Stack) GetHeight() float64 {
	log.Print("Stack::GetStackHeight")
	if s.Orientation == StackOrientationHorizontal {
		minY := math.MaxFloat64
		maxY := -math.MaxFloat64
		for _, child := range s.Children {
			minY = F64Min(minY, getMinY(child))
			maxY = F64Max(maxY, getMaxY(child))
		}
		return maxY - minY
	} else if s.Orientation == StackOrientationVertical {
		height := 0.0
		for _, child := range s.Children {
			height += getTotalHeight(child)
		}
		return height
	}

	panic("unknown stack orientation")
}

func F64Min(vars ...float64) float64 {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func F64Max(vars ...float64) float64 {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func getTotalWidth(e KeyboardElement) float64 {
	return e.GetWidth() + e.GetMargin()*2
}

func getTotalHeight(e KeyboardElement) float64 {
	return e.GetHeight() + e.GetMargin()*2
}

func getMinX(e KeyboardElement) float64 {
	return (-getTotalWidth(e) / 2) + e.GetXOffset()
}

func getMaxX(e KeyboardElement) float64 {
	return (getTotalWidth(e) / 2) + e.GetXOffset()
}

func getMinY(e KeyboardElement) float64 {
	return (-getTotalHeight(e) / 2) + e.GetYOffset()
}

func getMaxY(e KeyboardElement) float64 {
	return (getTotalHeight(e) / 2) + e.GetYOffset()
}
