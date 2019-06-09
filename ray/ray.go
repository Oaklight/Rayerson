package ray

import "vector"

type Ray struct {
	// A is the origin
	A *vector.Vec3
	// B is the direction
	B *vector.Vec3
}

func NewRay(a, b *vector.Vec3) *Ray {
	return &Ray{A: a, B: b}
}

func (r *Ray) Origin() *vector.Vec3 {
	return r.A
}

func (r *Ray) Direction() *vector.Vec3 {
	return r.B
}

func (r *Ray) PointAtPar(t float64) *vector.Vec3 {
	return vector.Add(r.A, vector.Mut(r.B, t))
}
