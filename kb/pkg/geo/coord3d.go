package geo

import "fmt"

// Coord3D represents a coordinate system in 3-dimensional space.
type Coord3D struct {
	Origin *Point3D
	Axes   *Matrix3D
	Name   string
}

func (x *Coord3D) String() string {
	panic("not yet implemented")
}

func NewCoord3D(name string) *Coord3D {
	if name == "" {
		name = generateCoordName()
	}

	return &Coord3D{
		Origin: &Point3D{X: 0, Y: 0, Z: 0},
		Axes:   identityMatrix,
		Name:   name,
	}
}

var (
	coordNameIndex = 0
)

func generateCoordName() string {
	name := fmt.Sprintf("Coord_%d", coordNameIndex)
	coordNameIndex++
	return name
}
