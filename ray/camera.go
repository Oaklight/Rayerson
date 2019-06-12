package ray

import (
	vec3 "vector"
)

// Camera defines a perspective camera model
type Camera struct {
	origin, lowerLeft, xAxis, yAxis *vec3.Vec3
}

// NewCamera Creates the default orthogonal camera model
func NewCamera() *Camera {
	return &Camera{
		origin:    vec3.Zeros(),
		lowerLeft: vec3.NewVec3(-2, -1, -1),
		xAxis:     vec3.NewVec3(4, 0, 0),
		yAxis:     vec3.NewVec3(0, 2, 0),
	}
}

// GetRay returns the ray at shifted NDC (u,v)
func (c *Camera) GetRay(u, v float64) *Ray {
	return NewRay(
		c.origin,
		vec3.Add(c.lowerLeft, c.xAxis.Mut(u), c.yAxis.Mut(v)).Sub(c.origin),
	)
}
