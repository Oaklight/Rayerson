package ray

import (
	"math"
	vec3 "vector"
)

// Camera defines a perspective camera model
type Camera struct {
	origin, lowerLeft, horizontal, vertical *vec3.Vec3
}

// NewCamera Creates the default orthogonal camera model
// ** lookAt is a direction
func NewCamera(fov, aspect float64, pos, lookAt, up vec3.Vec3) *Camera {
	theta := fov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := pos.Sub(&lookAt).Normalize()
	u := up.Cross(w).Normalize()
	v := w.Cross(u) // normalized already
	x := u.MulScalar(halfWidth)
	y := v.MulScalar(halfHeight)
	return &Camera{
		origin:     &pos,
		lowerLeft:  (&pos).Sub(x, y, w),
		horizontal: x.MulScalar(2),
		vertical:   y.MulScalar(2),
	}
}

// GetRay returns the ray at shifted NDC (u,v)
func (c *Camera) GetRay(u, v float64) *Ray {
	return NewRay(
		c.origin,
		vec3.Add(
			c.lowerLeft,
			c.horizontal.MulScalar(u),
			c.vertical.MulScalar(v),
			c.origin.Negate(),
		),
	)
}
