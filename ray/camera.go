package ray

import (
	"math"
	"math/rand"
	vec3 "vector"
)

// Camera defines a perspective camera model
type Camera struct {
	w, u, v                                 *vec3.Vec3
	origin, lowerLeft, horizontal, vertical *vec3.Vec3
	lensRadius                              float64
}

// NewCamera Creates the default orthogonal camera model
// ** lookAt is a direction
func NewCamera(fov, aspect, aperture float64, pos, lookAt, up vec3.Vec3) *Camera {
	theta := fov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := pos.Sub(&lookAt).Normalize()
	u := up.Cross(w).Normalize()
	v := w.Cross(u) // normalized already
	focusDist := (&pos).Sub(&lookAt).Length()
	x := u.MulScalar(halfWidth * focusDist)
	y := v.MulScalar(halfHeight * focusDist)
	return &Camera{
		origin:     &pos,
		lowerLeft:  (&pos).Sub(x, y, w.MulScalar(focusDist)),
		horizontal: x.MulScalar(2),
		vertical:   y.MulScalar(2),
		lensRadius: aperture / 2,
		w:          w,
		u:          u,
		v:          v,
	}
}

// GetRay returns the ray at shifted NDC (u,v)
func (c *Camera) GetRay(u, v float64, rnd *rand.Rand) *Ray {
	rd := randomInUnitDisc(rnd).MulScalar(c.lensRadius)
	offset := c.u.MulScalar(rd.X).Add(c.v.MulScalar(rd.Y))

	return NewRay(
		c.origin.Add(offset),
		vec3.Add(
			c.lowerLeft,
			c.horizontal.MulScalar(u),
			c.vertical.MulScalar(v),
		).Sub(
			c.origin,
			offset,
		),
	)
}

func randomInUnitDisc(rnd *rand.Rand) (rd *vec3.Vec3) {
	commonSub := &vec3.Vec3{1, 1, 0}
	for {
		rd1 := &vec3.Vec3{rnd.Float64(), rnd.Float64(), 0}
		rd = rd1.MulScalar(2).Sub(commonSub)
		if rd.Dot(rd) < 1.0 {
			return
		}
	}
}
