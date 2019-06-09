package ray

import "vector"

// Color is a type alias of vector.Vec3
type Color struct{ vector.Vec3 }

// R return the 1st channel of c
func (c *Color) R() float64 {
	return c.X()
}

// G returns the 2nd channel of c
func (c *Color) G() float64 {
	return c.Y()
}

// B returns the 3rd channel of c
func (c *Color) B() float64 {
	return c.Z()
}
