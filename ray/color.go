package ray

import "vector"

// Color64 is a 64-bit color array
type Color64 struct {
	R, G, B float64
}

// Vec2Color return a Color64 based on a given Vec3 object
func Vec2Color(v *vector.Vec3) *Color64 {
	return &Color64{v.X, v.Y, v.Z}
}

// NewColor64 returns a Color64 based on given rgb value, rgb should be int or int16
func NewColor64(r, g, b float64) *Color64 {
	return &Color64{r, g, b}
}

// Add use first argument as pivot vector, iterate to add rest of vectors
func Add(c *Color64, vs ...*Color64) *Color64 {
	e0, e1, e2 := c.R, c.G, c.B
	for _, v := range vs {
		e0 += v.R
		e1 += v.G
		e2 += v.B
	}
	sum := NewColor64(e0, e1, e2)
	return sum
}

// Add use first argument as pivot vector, iterate to add rest of vectors
func (c *Color64) Add(vs ...*Color64) *Color64 {
	e0, e1, e2 := c.R, c.G, c.B
	for _, v := range vs {
		e0 += v.R
		e1 += v.G
		e2 += v.B
	}
	sum := NewColor64(e0, e1, e2)
	return sum
}

// DivScalar performs scalar division on c
func DivScalar(c *Color64, s float64) *Color64 {
	return NewColor64(
		c.R/s,
		c.G/s,
		c.B/s,
	)
}

// DivScalar performs scalar division on c
func (c *Color64) DivScalar(s float64) *Color64 {
	return NewColor64(
		c.R/s,
		c.G/s,
		c.B/s,
	)
}
