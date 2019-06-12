package ray

import "vector"

// Ray comprises of an origin, and a direction
type Ray struct {
	Origin *vector.Vec3
	Direct *vector.Vec3
}

// NewRay creates and returns a Ray object from given values
func NewRay(a, b *vector.Vec3) *Ray {
	return &Ray{Origin: a, Direct: b}
}

// PointAtScale returns a point of t times the given Ray r, along its direction
func (r *Ray) PointAtScale(t float64) *vector.Vec3 {
	return vector.Add(r.Origin, vector.MulScalar(r.Direct, t))
}
