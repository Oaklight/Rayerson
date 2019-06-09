package ray

import "vector"

// Ray comprises of an origin A, and a direction B
type Ray struct {
	// A is the origin
	A *vector.Vec3
	// B is the direction
	B *vector.Vec3
}

// NewRay creates and returns a Ray object from given values
func NewRay(a, b *vector.Vec3) *Ray {
	return &Ray{A: a, B: b}
}

// Origin returns the origin of given Ray object
func (r *Ray) Origin() *vector.Vec3 {
	return r.A
}

// Direction returns the direction of given Ray object
func (r *Ray) Direction() *vector.Vec3 {
	return r.B
}

// PointAtScale returns a point of t times the given Ray r, along its direction
func (r *Ray) PointAtScale(t float64) *vector.Vec3 {
	return vector.Add(r.A, vector.Mut(r.B, t))
}
