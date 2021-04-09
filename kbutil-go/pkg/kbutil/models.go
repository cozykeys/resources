package kbutil

type ComponentBase struct {
	Name    string
	XOffset float64
	YOffset float64
}

type Component interface {
}

type Legend struct {
	Text string
}

type Key struct {
	ComponentBase

	Legends []Legend
}

type Circle struct {
	ComponentBase

	Radius float64
}

type Keyboard struct {
	Name       string
	Version    string
	Width      float64
	Height     float64
	Components []Component
}
