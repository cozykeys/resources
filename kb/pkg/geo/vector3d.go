package geo

import (
	"fmt"
	"math"
)

// Vector3D represents a vector in 3-dimensional space.
type Vector3D struct {
	X     float64
	Y     float64
	Z     float64
	Coord *Coord3D
}

func (v *Vector3D) String() string {
	return fmt.Sprintf("(%s, %s, %s)",
		formatFloat(v.X),
		formatFloat(v.Y),
		formatFloat(v.Z))
}

func NewVector3DFromPoints(p1, p2 *Point3D) *Vector3D {
	if p1.Coord != p2.Coord {
		p2 = p2.ConvertTo(p1.Coord)
	}

	return &Vector3D{
		X:     p2.X - p1.X,
		Y:     p2.Y - p1.Y,
		Z:     p2.Z - p1.Z,
		Coord: p1.Coord,
	}
}

func (v *Vector3D) GetDirection() *Vector3D {
	return v
}

func (v *Vector3D) Multiply(distance float64) *Vector3D {
	return &Vector3D{
		X: v.X * distance,
		Y: v.Y * distance,
		Z: v.Z * distance,
	}
}

func (v *Vector3D) Normalized() *Vector3D {
	norm := v.Norm()
	return &Vector3D{
		X: v.X / norm,
		Y: v.Y / norm,
		Z: v.Z / norm,
	}

}

func (v *Vector3D) Norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3D) IsParallelTo(obj LinearObject) bool {
	tmp := obj.GetDirection()
	if v.Coord != tmp.Coord {
		tmp = tmp.ConvertTo(v.Coord)
	}

	return equal(v.Normalized().Cross(v.Normalized()).Norm(), 0.0)
}

func (v *Vector3D) ConvertTo(coord *Coord3D) *Vector3D {
	// TODO: Not yet implemented
	return v
}

func (v *Vector3D) Cross(other *Vector3D) *Vector3D {
	if v.Coord != other.Coord {
		other = other.ConvertTo(v.Coord)
	}

	return &Vector3D{
		X:     v.Y*other.Z - v.Z*other.Y,
		Y:     v.Z*other.X - v.X*other.Z,
		Z:     v.X*other.Y - v.Y*other.X,
		Coord: v.Coord,
	}
}

func (v *Vector3D) ToPoint() *Point3D {
	return &Point3D{
		X:     v.X,
		Y:     v.Y,
		Z:     v.Z,
		Coord: v.Coord,
	}
}

func (v *Vector3D) Subtract(other *Vector3D) *Vector3D {
	if v.Coord != other.Coord {
		other = other.ConvertTo(v.Coord)
	}

	return &Vector3D{
		X:     v.X - other.X,
		Y:     v.Y - other.Y,
		Z:     v.Z - other.Z,
		Coord: v.Coord,
	}
}

func (v *Vector3D) Add(other *Vector3D) *Vector3D {
	if v.Coord != other.Coord {
		other = other.ConvertTo(v.Coord)
	}

	return &Vector3D{
		X:     v.X + other.X,
		Y:     v.Y + other.Y,
		Z:     v.Z + other.Z,
		Coord: v.Coord,
	}
}

func (v *Vector3D) Mult(a float64) *Vector3D {
	return &Vector3D{
		X:     v.X * a,
		Y:     v.Y * a,
		Z:     v.Z * a,
		Coord: v.Coord,
	}
}

func (v *Vector3D) Dot(other *Vector3D) float64 {
	if v.Coord != other.Coord {
		other = other.ConvertTo(v.Coord)
	}

	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}
