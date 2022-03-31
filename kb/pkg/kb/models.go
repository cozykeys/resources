package kb

// Circle TODO
type Circle struct {
	Name    string
	Radius  float64
	XOffset float64
	YOffset float64
}

// Keyboard TODO
type Keyboard struct {
	Name       string
	Version    string
	Width      float64
	Height     float64
	Components []Component
}

// Key TODO
type Key struct {
	Name    string
	Legends []Legend
	XOffset float64
	YOffset float64
}

// Legend TODO
type Legend struct {
	Text string
}

// Component TODO
type Component interface{}
