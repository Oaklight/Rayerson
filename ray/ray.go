package ray

import "vector"

// Ray comprises of an origin, and a direction
type Ray struct {
	origin *vector.Vec3
	direct *vector.Vec3
}

// NewRay creates and returns a Ray object from given values
func NewRay(a, b *vector.Vec3) *Ray {
	return &Ray{origin: a, direct: b}
}

// Origin returns the origin of given Ray object
func (r *Ray) Origin() *vector.Vec3 {
	return r.origin
}

// Direction returns the direction of given Ray object
func (r *Ray) Direction() *vector.Vec3 {
	return r.direct
}

// PointAtScale returns a point of t times the given Ray r, along its direction
func (r *Ray) PointAtScale(t float64) *vector.Vec3 {
	return vector.Add(r.origin, vector.Mut(r.direct, t))
}
