package geo

type LinearObject interface {
	GetDirection() *Vector3D
}

type FiniteObject interface {
	PointLocation(*Point3D) int
}
