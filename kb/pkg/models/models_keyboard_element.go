package models

// KeyboardElement TODO
type KeyboardElement interface {
	GetName() string
	GetXOffset() float64
	GetYOffset() float64
	GetRotation() float64
	GetHeight() float64
	GetWidth() float64
	GetMargin() float64
	GetDebug() bool
	GetVisible() bool
	GetParent() KeyboardElement
	GetConstants() map[string]string
}

// KeyboardElementBase TODO
type KeyboardElementBase struct {
	Name      string
	XOffset   float64
	YOffset   float64
	Rotation  float64
	Height    float64
	Width     float64
	Margin    float64
	Debug     bool
	Visible   bool
	Parent    KeyboardElement
	Constants map[string]string
}

// GetName TODO
func (e *KeyboardElementBase) GetName() string {
	return e.Name
}

// GetXOffset TODO
func (e *KeyboardElementBase) GetXOffset() float64 {
	return e.XOffset
}

// GetYOffset TODO
func (e *KeyboardElementBase) GetYOffset() float64 {
	return e.YOffset
}

// GetRotation TODO
func (e *KeyboardElementBase) GetRotation() float64 {
	return e.Rotation
}

// GetHeight TODO
func (e *KeyboardElementBase) GetHeight() float64 {
	return e.Height
}

// GetWidth TODO
func (e *KeyboardElementBase) GetWidth() float64 {
	return e.Width
}

// GetMargin TODO
func (e *KeyboardElementBase) GetMargin() float64 {
	return e.Margin
}

// GetDebug TODO
func (e *KeyboardElementBase) GetDebug() bool {
	return e.Debug
}

// GetVisible TODO
func (e *KeyboardElementBase) GetVisible() bool {
	return e.Visible
}

// GetParent TODO
func (e *KeyboardElementBase) GetParent() KeyboardElement {
	return e.Parent
}

// GetConstants TODO
func (e *KeyboardElementBase) GetConstants() map[string]string {
	constants := e.Constants
	parent := e.GetParent()
	for parent != nil {
		parentConstants := parent.GetConstants()
		constants = mergeConstants(parentConstants, constants)
		parent = parent.GetParent()
	}
	return constants
}

var _ KeyboardElement = &KeyboardElementBase{}
