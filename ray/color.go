package ray

import (
	"image/color"
	"math"
	"vector"
)

// Color is a 64-bit color array
type Color struct {
	R, G, B float64
}

var (
	Black = Color{0, 0, 0}
	White = Color{255, 255, 255}
)

// Vec2Color return a Color based on a given Vec3 object
func Vec2Color(v *vector.Vec3) *Color {
	return &Color{v.X, v.Y, v.Z}
}

// NewColor returns a Color based on given rgb value, rgb should be int or int16
func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

// =============================Vec3 Class methods=============================

// RGBA for compatibility with image.Color
func (c *Color) RGBA() color.RGBA {
	r := uint8(math.Max(0, math.Min(255, c.R*255)))
	g := uint8(math.Max(0, math.Min(255, c.G*255)))
	b := uint8(math.Max(0, math.Min(255, c.B*255)))
	return color.RGBA{r, g, b, 255}
}

// RGBA64 for compatibility with image.Color
func (c *Color) RGBA64() color.RGBA64 {
	r := uint16(math.Max(0, math.Min(65535, c.R*65535)))
	g := uint16(math.Max(0, math.Min(65535, c.G*65535)))
	b := uint16(math.Max(0, math.Min(65535, c.B*65535)))
	return color.RGBA64{r, g, b, 65535}
}

// Add use first argument as pivot vector, iterate to add rest of vectors
func (c *Color) Add(cs ...*Color) *Color {
	e0, e1, e2 := c.R, c.G, c.B
	for _, each := range cs {
		e0 += each.R
		e1 += each.G
		e2 += each.B
	}
	return &Color{e0, e1, e2}
}

// Mul should only be used for attenuation, which is a scaling factor vector
func (c *Color) Mul(s *Color) *Color {
	return &Color{c.R * s.R, c.G * s.G, c.B * s.B}
}

// DivScalar performs scalar division on c
func (c *Color) DivScalar(s float64) *Color {
	return &Color{
		c.R / s,
		c.G / s,
		c.B / s,
	}
}

// MulScalar performs scalar division on c
func (c *Color) MulScalar(s float64) *Color {
	return &Color{
		c.R * s,
		c.G * s,
		c.B * s,
	}
}

// =============================General function===============================

// Add use first argument as pivot vector, iterate to add rest of vectors
func Add(c *Color, vs ...*Color) *Color {
	return c.Add(vs...)
}

// Mul should only be used for attenuation, which is a scaling factor vector
func Mul(c *Color, s *Color) *Color {
	return c.Mul(s)
}

// DivScalar performs scalar division on c
func DivScalar(c *Color, s float64) *Color {
	return c.DivScalar(s)
}

// MulScalar performs scalar division on c
func MulScalar(c *Color, s float64) *Color {
	return c.MulScalar(s)
}
