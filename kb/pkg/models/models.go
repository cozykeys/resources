package models

// Constant TODO
type Constant struct {
	Name  string
	Value string
}

// Keyboard TODO
type Keyboard struct {
	//ComponentBase

	Name      string
	Version   string
	Width     float64
	Height    float64
	Constants []Constant
	Layers    []Layer
}

var _ Component = &Keyboard{}

// ComponentBase TODO
type ComponentBase struct {
	Name      string
	XOffset   float64
	YOffset   float64
	Rotation  float64
	Height    float64
	Width     float64
	Margin    float64
	Parent    Component
	Constants map[string]Constant
	Debug     bool
	Visible   bool
}

type Component interface{}

// Circle TODO
type Circle struct {
	ComponentBase

	Size        float64
	Fill        string
	Stroke      string
	StrokeWidth string
}

var _ Component = &Circle{}

// Key TODO
type Key struct {
	ComponentBase

	Legends []Legend
	Row     int
	Column  int
	Fill    string
	Stroke  string
}

var _ Component = &Key{}

// Legend TODO
type Legend struct {
	ComponentBase

	Text                string
	FontSize            float64
	HorizontalAlignment LegendHorizontalAlignment
	VerticalAlignment   LegendVerticalAlignment
	Color               string
}

var _ Component = &Legend{}

// Group TODO
type Group struct {
	ComponentBase

	Children []Component
}

var _ Component = &Group{}

// Layer TODO
type Layer struct {
	ComponentBase

	ZIndex int
	Groups []Group
}

var _ Component = &Layer{}

// Spacer TODO
type Spacer struct {
	ComponentBase
}

var _ Component = &Spacer{}

// Stack TODO
type Stack struct {
	Group

	Orientation StackOrientation
}

var _ Component = &Stack{}

// Text TODO
type Text struct {
	Content    string
	TextAnchor string
	Font       string
	Fill       string
}

var _ Component = &Text{}

// LegendHorizontalAlignment TODO
type LegendHorizontalAlignment int

const (
	LegendHorizontalAlignmentLeft LegendHorizontalAlignment = iota
	LegendHorizontalAlignmentCenter
	LegendHorizontalAlignmentRight
)

// LegendVerticalAlignment TODO
type LegendVerticalAlignment int

const (
	LegendVerticalAlignmentTop LegendVerticalAlignment = iota
	LegendVerticalAlignmentCenter
	LegendVerticalAlignmentBottom
)

// StackOrientation TODO
type StackOrientation int

const (
	StackOrientationHorizontal StackOrientation = iota
	StackOrientationVertical
)

type PathComponent interface {
	Data() string
}

// AbsoluteCubicCurveTo TODO
type AbsoluteCubicCurveTo struct {
	ComponentBase

	EndPoint      Vec2
	ControlPointA Vec2
	ControlPointB Vec2
}

var _ Component = &AbsoluteCubicCurveTo{}
var _ PathComponent = &AbsoluteCubicCurveTo{}

func (x *AbsoluteCubicCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// AbsoluteHorizontalLineTo TODO
type AbsoluteHorizontalLineTo struct {
	ComponentBase

	X float64
}

var _ Component = &AbsoluteHorizontalLineTo{}
var _ PathComponent = &AbsoluteHorizontalLineTo{}

func (x *AbsoluteHorizontalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// AbsoluteLineTo TODO
type AbsoluteLineTo struct {
	ComponentBase

	EndPoint Vec2
}

var _ Component = &AbsoluteLineTo{}
var _ PathComponent = &AbsoluteLineTo{}

func (x *AbsoluteLineTo) Data() string {
	//string Data => $"L {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}

// AbsoluteMoveTo TODO
type AbsoluteMoveTo struct {
	ComponentBase

	EndPoint Vec2
}

var _ Component = &AbsoluteMoveTo{}
var _ PathComponent = &AbsoluteMoveTo{}

func (x *AbsoluteMoveTo) Data() string {
	//string Data => $"M {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}

// AbsoluteQuadraticCurveTo TODO
type AbsoluteQuadraticCurveTo struct {
	ComponentBase

	EndPoint     Vec2
	ControlPoint Vec2
}

var _ Component = &AbsoluteQuadraticCurveTo{}
var _ PathComponent = &AbsoluteQuadraticCurveTo{}

func (x *AbsoluteQuadraticCurveTo) Data() string {
	//string Data => $"Q {ControlPoint.X} {ControlPoint.Y} {EndPoint.X} {EndPoint.Y}";
	return "NYI"
}

// AbsoluteVerticalLineTo TODO
type AbsoluteVerticalLineTo struct {
	ComponentBase

	Y float64
}

var _ Component = &AbsoluteVerticalLineTo{}
var _ PathComponent = &AbsoluteVerticalLineTo{}

func (x *AbsoluteVerticalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// Path TODO
type Path struct //
{
	ComponentBase

	Fill        string
	FillOpacity string
	Stroke      string
	StrokeWidth string
	Components  []PathComponent
}

var _ Component = &Path{}

func (x *Path) Data() string {
	//string Data => string.Join(" ", Components.Select(component => component.Data));
	return "NYI"
}

// RelativeCubicCurveTo TODO
type RelativeCubicCurveTo struct {
	ComponentBase

	EndPoint      Vec2
	ControlPointA Vec2
	ControlPointB Vec2
}

var _ Component = &RelativeCubicCurveTo{}
var _ PathComponent = &RelativeCubicCurveTo{}

func (x *RelativeCubicCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeHorizontalLineTo TODO
type RelativeHorizontalLineTo struct {
	ComponentBase

	X float64
}

var _ Component = &RelativeHorizontalLineTo{}
var _ PathComponent = &RelativeHorizontalLineTo{}

func (x *RelativeHorizontalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeLineTo TODO
type RelativeLineTo struct {
	ComponentBase

	EndPoint Vec2
}

var _ Component = &RelativeLineTo{}
var _ PathComponent = &RelativeLineTo{}

func (x *RelativeLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeMoveTo TODO
type RelativeMoveTo struct {
	ComponentBase

	EndPoint Vec2
}

var _ Component = &RelativeMoveTo{}
var _ PathComponent = &RelativeMoveTo{}

func (x *RelativeMoveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeQuadraticCurveTo TODO
type RelativeQuadraticCurveTo struct {
	ComponentBase

	EndPoint     Vec2
	ControlPoint Vec2
}

var _ Component = &RelativeQuadraticCurveTo{}
var _ PathComponent = &RelativeQuadraticCurveTo{}

func (x *RelativeQuadraticCurveTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

// RelativeVerticalLineTo TODO
type RelativeVerticalLineTo struct {
	ComponentBase

	Y float64
}

var _ Component = &RelativeVerticalLineTo{}
var _ PathComponent = &RelativeVerticalLineTo{}

func (x *RelativeVerticalLineTo) Data() string {
	//string Data => throw new System.NotImplementedException();
	return "NYI"
}

type Vec2 struct {
	ComponentBase

	X float64
	Y float64
}

var _ Component = &Vec2{}
