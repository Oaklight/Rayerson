package vector

import (
	"math"
)

// Vec3 defines the 3-dimensional vector
type Vec3 struct {
	X, Y, Z float64
}

var (
	Zeros = Vec3{}
	Ones  = Vec3{1, 1, 1}
)

// NewVec3 creates and returns a new vector ptr with given values
func NewVec3(e0, e1, e2 float64) *Vec3 {
	return &Vec3{e0, e1, e2}
}

// =============================Vec3 Class methods=============================

// Length returns the (square root) length of v1
func (v1 *Vec3) Length() float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
}

// LengthN returns the n-th length of v1
func (v1 *Vec3) LengthN(n float64) float64 {
	if n == 2 {
		return v1.Length()
	}
	v1 = v1.Abs()
	return math.Pow(math.Pow(v1.X, n)+math.Pow(v1.Y, n)+math.Pow(v1.Z, n), 1/n)
}

// Dot performs dot product, v1 as the pivot vector
func (v1 *Vec3) Dot(v2 *Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Cross performs the cross product, v1 as the pivot vector
func (v1 *Vec3) Cross(v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Z*v2.X,
	)
}

// using variadics to support multi-variable Add, Sub

// Add use first argument as pivot vector, iterate to add rest of vectors
func (v1 *Vec3) Add(vs ...*Vec3) *Vec3 {
	e0, e1, e2 := v1.X, v1.Y, v1.Z
	for _, v := range vs {
		e0 += v.X
		e1 += v.Y
		e2 += v.Z
	}
	sum := NewVec3(e0, e1, e2)
	return sum
}

// Sub use first argument as pivot vector, iterate to substract rest of vectors
func (v1 *Vec3) Sub(vs ...*Vec3) *Vec3 {
	e0, e1, e2 := v1.X, v1.Y, v1.Z
	for _, v := range vs {
		e0 -= v.X
		e1 -= v.Y
		e2 -= v.Z
	}
	res := NewVec3(e0, e1, e2)
	return res
}

// MulScalar performs scalar multiplication on v1
func (v1 *Vec3) MulScalar(s float64) *Vec3 {
	return NewVec3(
		v1.X*s,
		v1.Y*s,
		v1.Z*s,
	)
}

// DivScalar performs scalar division on v1
func (v1 *Vec3) DivScalar(s float64) *Vec3 {
	return NewVec3(
		v1.X/s,
		v1.Y/s,
		v1.Z/s,
	)
}

// Normalize return the unit-length vector along v1 direction
func (v1 *Vec3) Normalize() *Vec3 {
	return v1.DivScalar(v1.Length())
}

// Abs returns the absoluted value of original vector
func (v1 *Vec3) Abs() *Vec3 {
	return &Vec3{math.Abs(v1.X), math.Abs(v1.Y), math.Abs(v1.Z)}
}

// =============================General function===============================

// Length returns the (square root) length of v1
func Length(v1 *Vec3) float64 {
	return v1.Length()
}

// LengthN returns the n-th length of v1
func LengthN(v1 *Vec3, n float64) float64 {
	return v1.LengthN(n)
}

// Dot performs dot product, v1 as the pivot vector
func Dot(v1, v2 *Vec3) float64 {
	return v1.Dot(v2)
}

// Cross performs the cross product, v1 as the pivot vector
func Cross(v1, v2 *Vec3) *Vec3 {
	return v1.Cross(v2)
}

// Add use first argument as pivot vector, iterate to add rest of vectors
func Add(v1 *Vec3, vs ...*Vec3) *Vec3 {
	return v1.Add(vs...)
}

// Sub use first argument as pivot vector, iterate to substract rest of vectors
func Sub(v1 *Vec3, vs ...*Vec3) *Vec3 {
	return v1.Sub(vs...)
}

// MulScalar performs scalar multiplication on v1
func MulScalar(v1 *Vec3, s float64) *Vec3 {
	return v1.MulScalar(s)
}

// DivScalar performs scalar division on v1
func DivScalar(v1 *Vec3, s float64) *Vec3 {
	return v1.DivScalar(s)
}

// Normalize return the unit-length vector along v1 direction
func Normalize(v1 *Vec3) *Vec3 {
	return v1.Normalize()
}

// Abs returns the absoluted value of original vector
func Abs(v1 *Vec3) *Vec3 {
	return v1.Abs()
}
