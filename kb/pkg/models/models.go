package models

type KeyboardElement interface {
	GetParent() KeyboardElement
	GetConstants() []Constant
}

type KeyboardElementBase struct {
	Parent    KeyboardElement
	Constants []Constant
}

func (e *KeyboardElementBase) GetParent() KeyboardElement {
	return e.Parent
}

func (e *KeyboardElementBase) GetConstants() []Constant {
	constants := e.Constants
	parent := e.GetParent()
	for parent != nil {
		parentConstants := parent.GetConstants()
		constants = mergeConstants(parentConstants, constants)
		parent = parent.GetParent()
	}
	return constants
}

// Constant TODO
type Constant struct {
	KeyboardElementBase

	Name  string
	Value string
}

// Keyboard TODO
type Keyboard struct {
	KeyboardElementBase

	Name    string
	Version string
	Width   float64
	Height  float64
	Layers  []Layer
}

// Circle TODO
type Circle struct {
	KeyboardElementBase

	Size        float64
	Fill        string
	Stroke      string
	StrokeWidth string
	XOffset     float64
	YOffset     float64
	Name        string
}

// Key TODO
type Key struct {
	KeyboardElementBase

	Legends []Legend
	Row     int
	Column  int
	Fill    string
	Stroke  string
	XOffset float64
	YOffset float64
	Name    string
	Width   float64
	Height  float64
	Margin  float64
}

// Legend TODO
type Legend struct {
	KeyboardElementBase

	Text                string
	FontSize            float64
	HorizontalAlignment LegendHorizontalAlignment
	VerticalAlignment   LegendVerticalAlignment
	Color               string
	YOffset             float64
}

type GroupChild interface{}

// Group TODO
type Group struct {
	KeyboardElementBase

	Name     string
	Rotation float64
	XOffset  float64
	YOffset  float64
	Visible  bool

	Children []GroupChild
}

// Layer TODO
type Layer struct {
	KeyboardElementBase

	ZIndex  int
	Groups  []Group
	XOffset float64
	YOffset float64
	Name    string
}

// Spacer TODO
type Spacer struct {
	KeyboardElementBase

	Height float64
	Width  float64
}

// Stack TODO
type Stack struct {
	Group

	Orientation StackOrientation
}

// Text TODO
type Text struct {
	KeyboardElementBase

	Content    string
	TextAnchor string
	Font       string
	Fill       string
	XOffset    float64
	YOffset    float64
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

type PathComponent interface {
	Data() string
}

// AbsoluteCubicCurveTo TODO
type AbsoluteCubicCurveTo struct {
	KeyboardElementBase

	EndPoint      Point
	ControlPointA Point
	ControlPointB Point
}

var _ PathComponent = &AbsoluteCubicCurveTo{}

func (x *AbsoluteCubicCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// AbsoluteHorizontalLineTo TODO
type AbsoluteHorizontalLineTo struct {
	KeyboardElementBase

	X float64
}

var _ PathComponent = &AbsoluteHorizontalLineTo{}

func (x *AbsoluteHorizontalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// AbsoluteLineTo TODO
type AbsoluteLineTo struct {
	KeyboardElementBase

	EndPoint *Point
}

var _ PathComponent = &AbsoluteLineTo{}

func (x *AbsoluteLineTo) Data() string {
	//string Data => $"L {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}

// AbsoluteMoveTo TODO
type AbsoluteMoveTo struct {
	KeyboardElementBase

	EndPoint *Point
}

var _ PathComponent = &AbsoluteMoveTo{}

func (x *AbsoluteMoveTo) Data() string {
	//string Data => $"M {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}

// AbsoluteQuadraticCurveTo TODO
type AbsoluteQuadraticCurveTo struct {
	KeyboardElementBase

	EndPoint     Point
	ControlPoint Point
}

var _ PathComponent = &AbsoluteQuadraticCurveTo{}

func (x *AbsoluteQuadraticCurveTo) Data() string {
	//string Data => $"Q {ControlPoint.X} {ControlPoint.Y} {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}

// AbsoluteVerticalLineTo TODO
type AbsoluteVerticalLineTo struct {
	KeyboardElementBase

	Y float64
}

var _ PathComponent = &AbsoluteVerticalLineTo{}

func (x *AbsoluteVerticalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// Path TODO
type Path struct {
	KeyboardElementBase

	Fill        string
	FillOpacity string
	Stroke      string
	StrokeWidth string
	Visible     bool
	Components  []PathComponent
}

func (x *Path) Data() string {
	//string Data => string.Join(" ", Components.Select(component => component.Data));
	return "NYI"
}

// RelativeCubicCurveTo TODO
type RelativeCubicCurveTo struct {
	KeyboardElementBase

	EndPoint      Point
	ControlPointA Point
	ControlPointB Point
}

var _ PathComponent = &RelativeCubicCurveTo{}

func (x *RelativeCubicCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeHorizontalLineTo TODO
type RelativeHorizontalLineTo struct {
	KeyboardElementBase

	X float64
}

var _ PathComponent = &RelativeHorizontalLineTo{}

func (x *RelativeHorizontalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeLineTo TODO
type RelativeLineTo struct {
	KeyboardElementBase

	EndPoint Point
}

var _ PathComponent = &RelativeLineTo{}

func (x *RelativeLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeMoveTo TODO
type RelativeMoveTo struct {
	KeyboardElementBase

	EndPoint Point
}

var _ PathComponent = &RelativeMoveTo{}

func (x *RelativeMoveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeQuadraticCurveTo TODO
type RelativeQuadraticCurveTo struct {
	KeyboardElementBase

	EndPoint     Point
	ControlPoint Point
}

var _ PathComponent = &RelativeQuadraticCurveTo{}

func (x *RelativeQuadraticCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeVerticalLineTo TODO
type RelativeVerticalLineTo struct {
	KeyboardElementBase

	Y float64
}

var _ PathComponent = &RelativeVerticalLineTo{}

func (x *RelativeVerticalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

type Point struct {
	KeyboardElementBase

	X float64
	Y float64
}
