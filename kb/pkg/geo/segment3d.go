package geo

import (
	"fmt"
	"math"
)

// Segment3D represents a segment in 3-dimensional space.
type Segment3D struct {
	P1 *Point3D
	P2 *Point3D
}

func (s *Segment3D) String() string {
	return fmt.Sprintf("[%v, %v]", s.P1, s.P2)
}

func (s *Segment3D) ToLine() *Line3D {
	return NewLine3DFromPoints(s.P1, s.P2)
}

// TODO: This name is terrible; think of one that is more descriptive
func (s *Segment3D) GetPoint(distance float64) *Point3D {
	if distance > s.Magnitude() {
		panic("TODO")
	}

	theta := s.Theta()

	dx := distance * math.Cos(theta)
	dy := distance * math.Sin(theta)

	return &Point3D{
		X: s.P1.X + dx,
		Y: s.P1.Y + dy,
	}
}

func (s *Segment3D) Magnitude() float64 {
	return s.P1.DistanceToPoint(s.P2)
}

func (s *Segment3D) GetDeltaX() float64 {
	return s.P2.X - s.P1.X
}

func (s *Segment3D) GetDeltaY() float64 {
	return s.P2.Y - s.P1.Y
}

func (s *Segment3D) GetDeltaZ() float64 {
	return s.P2.Z - s.P1.Z
}

// TODO: Does this make sense on a 3D Segment? Should we assert that Z == 0?
// Maybe rename to Theta2D?
func (s *Segment3D) Theta() float64 {
	dx := s.GetDeltaX()
	dy := s.GetDeltaY()

	// Early out if the segment is parallel to an axis.
	// TODO: Make and use "greater"/"lesser" functions that uses the same
	// tolerance as equal
	if equal(dy, 0) && dx > 0 {
		return 0.0 * math.Pi
	}
	if equal(dx, 0) && dy > 0 {
		return 0.5 * math.Pi
	}
	if equal(dy, 0) && dx < 0 {
		return 1.0 * math.Pi
	}
	if equal(dx, 0) && dy < 0 {
		return 1.5 * math.Pi
	}

	////double theta = Math.Atan(dy / dx);
	theta := math.Atan(dy / dx)

	// Because slope doesn't take direction into account, we have to manually
	// adjust theta for segments whose direction points into quadrants 2 and 3.
	if (dx < 0 && dy > 0) || (dx < 0 && dy < 0) {
		theta += math.Pi
	}

	return float64Mod(theta, math.Pi*2)
}

// TODO: Name this better; move to a Segment2D struct
func (s *Segment3D) PerpendicularClockwise(distance float64) *Segment3D {
	dx := s.P2.X - s.P1.X
	dy := s.P2.Y - s.P1.Y
	v := (&Vector3D{X: dy, Y: -dx}).Normalized().Multiply(distance)
	return &Segment3D{
		P1: s.P1.Translate(v),
		P2: s.P2.Translate(v),
	}
}

// TODO: Name this better; move to a Segment2D struct
func (s *Segment3D) PerpendicularCounterClockwise(distance float64) Segment3D {
	dx := s.P2.X - s.P1.X
	dy := s.P2.Y - s.P1.Y
	v := (&Vector3D{X: -dy, Y: dx}).Normalized().Multiply(distance)
	return Segment3D{
		P1: s.P1.Translate(v),
		P2: s.P2.Translate(v),
	}
}
