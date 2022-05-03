package geo

import "fmt"

// Line3D represents a line in 3-dimensional space.
type Line3D struct {
	Point     *Point3D
	Direction *Vector3D
}

var _ FiniteObject = &Line3D{}
var _ LinearObject = &Line3D{}

func (l *Line3D) GetPoint() *Point3D {
	if l == nil {
		return nil
	}

	return l.Point
}

func (l *Line3D) GetDirection() *Vector3D {
	if l == nil {
		return nil
	}

	return l.Direction
}

func (l *Line3D) String() string {
	return fmt.Sprintf("{ Point: %v, Dir: %v }", l.GetPoint(), l.GetDirection())
}

func (l *Line3D) PointLocation(p *Point3D) int {
	panic("not yet implemented")
}

// NewLine3DFromPoints initializes a line from two points.
func NewLine3DFromPoints(p1, p2 *Point3D) *Line3D {
	return &Line3D{
		Point:     p1,
		Direction: NewVector3DFromPoints(p1, p2),
	}
}

func (l *Line3D) IntersectionWith(other *Line3D) interface{} {
	if other.IsParallelTo(l) && other.Point.BelongsTo(l) {
		return l
	}

	p := other.PerpendicularTo(l)
	if p.BelongsTo(other) {
		return p
	} else {
		return nil
	}
}

func (l *Line3D) IsParallelTo(obj LinearObject) bool {
	d := obj.GetDirection()
	return l.Direction.IsParallelTo(d)
}

func (l *Line3D) PerpendicularTo(other *Line3D) *Point3D {
	r1 := l.Point.ToVector()
	r2 := other.Point.ToVector()
	s1 := l.Direction
	s2 := other.Direction
	if s1.Cross(s2).Norm() > tolerance {
		a := r2.Subtract(r1)
		b := a.Dot(s1.Cross(s1.Cross(s2)))
		c := s1.Dot(s2.Cross(s1.Cross(s2)))
		d := b / c
		r1 = r2.Add(s2.Mult(d))
		return r1.ToPoint()
	} else {
		// Lines are parallel; should we just return any point on the line? Or error?
		return other.Point
		//panic("Lines are parallel")
	}
}
