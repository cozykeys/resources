package geo

import "fmt"

var (
	identityMatrix = &Matrix3D{
		Data: [][]float64{
			{1.0, 0.0, 0.0},
			{0.0, 1.0, 0.0},
			{0.0, 0.0, 1.0},
		},
	}
)

// Matrix3D represents a matrix in 3-dimensional space.
type Matrix3D struct {
	Data [][]float64
}

func (m *Matrix3D) String() string {
	return fmt.Sprintf("[[%f, %f, %f], [%f, %f, %f], [%f, %f, %f]]",
		m.Data[0][0], m.Data[0][1], m.Data[0][2],
		m.Data[1][0], m.Data[1][1], m.Data[1][2],
		m.Data[2][0], m.Data[2][1], m.Data[2][2])
}
