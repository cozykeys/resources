package geo

import (
	"fmt"
	"math"
)

type Point3D struct {
	X     float64
	Y     float64
	Z     float64
	Coord *Coord3D
}

func (p *Point3D) String() string {
	return fmt.Sprintf("(%s, %s, %s)",
		formatFloat(p.X),
		formatFloat(p.Y),
		formatFloat(p.Z))
}

func (p *Point3D) Translate(v *Vector3D) *Point3D {
	return &Point3D{
		X: p.X + v.X,
		Y: p.Y + v.Y,
		Z: p.Z + v.Z,
	}
}

func (p *Point3D) ConvertTo(coord *Coord3D) *Point3D {
	// TODO: Not yet implemented
	return p
}

func (p *Point3D) BelongsTo(l *Line3D) bool {
	return p.DistanceToLine(l) <= tolerance
}

func (p *Point3D) ToVector() *Vector3D {
	return &Vector3D{
		X:     p.X,
		Y:     p.Y,
		Z:     p.Z,
		Coord: p.Coord,
	}
}

func (p *Point3D) DistanceToLine(l *Line3D) float64 {
	v := NewVector3DFromPoints(p, l.Point)
	return v.Cross(l.Direction).Norm() / l.Direction.Norm()
}

func (p *Point3D) DistanceToPoint(p2 *Point3D) float64 {
	if p.Coord != p2.Coord {
		p2 = p2.ConvertTo(p.Coord)
	}

	return math.Sqrt(p.X - p2.X*p.X - p2.X + p.Y - p2.Y*p.Y - p2.Y + p.Z - p2.Z*p.Z - p2.Z)
}
