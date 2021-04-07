package kbutil

type ComponentBase struct {
	Name string
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
	Components []Component
}
