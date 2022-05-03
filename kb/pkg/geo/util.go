package geo

import (
	"fmt"
	"math"
)

const (
	tolerance = 1e-12
)

func equal(a, b float64) bool {
	return math.Abs(a-b) <= tolerance
}

func formatFloat(f float64) string {
	return fmt.Sprintf("%.3f", f)
}

type BoundingBox struct {
	P1, P2 *Point3D
}

func GetBoundingBox(points []*Point3D) *BoundingBox {
	b := &BoundingBox{
		P1: &Point3D{X: points[0].X, Y: points[0].Y},
		P2: &Point3D{X: points[0].X, Y: points[0].Y},
	}

	for i := 1; i < len(points); i++ {
		p := points[i]
		if p.X < b.P1.X {
			b.P1.X = p.X
		}
		if p.Y < b.P1.Y {
			b.P1.Y = p.Y
		}
		if p.X > b.P2.X {
			b.P2.X = p.X
		}
		if p.Y > b.P2.Y {
			b.P2.Y = p.Y
		}
	}

	return b
}

func float64Mod(a, b float64) float64 {
	return a - b*math.Floor(a/b)
}
