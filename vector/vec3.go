package vector

import "math"

// Vec3 defines the 3-dimensional vector
type Vec3 struct {
	e [3]float64
}

// NewVec3 creates and returns a new vector ptr with given values
func NewVec3(e0, e1, e2 float64) *Vec3 {
	var v1 Vec3
	v1.e[0] = e0
	v1.e[1] = e1
	v1.e[2] = e2
	return &v1
}

// Zeros creates all zero vector
func Zeros() *Vec3 {
	return NewVec3(0, 0, 0)
}

// Ones creates all one vector
func Ones() *Vec3 {
	return NewVec3(1, 1, 1)
}

// X returns the first element of v1
func (v1 *Vec3) X() float64 {
	return v1.e[0]
}

// Y returns the second element of v1
func (v1 *Vec3) Y() float64 {
	return v1.e[1]
}

// Z returns the third element of v1
func (v1 *Vec3) Z() float64 {
	return v1.e[2]
}

// At returns the i-th element of v1, panic when index out of range
func (v1 *Vec3) At(i int) float64 {
	switch i {
	case 0, 1, 2:
		return v1.e[i]
	}
	panic("Index out of range inside Vec3.At()")
}

// Length returns the (square root) length of v1
func Length(v1 *Vec3) float64 {
	return math.Sqrt(v1.X()*v1.X() + v1.Y()*v1.Y() + v1.Z()*v1.Z())
}

// Length returns the (square root) length of v1
func (v1 *Vec3) Length() float64 {
	return math.Sqrt(v1.X()*v1.X() + v1.Y()*v1.Y() + v1.Z()*v1.Z())
}

// SqrLength returns the squared length of v1
func SqrLength(v1 *Vec3) float64 {
	return v1.X()*v1.X() + v1.Y()*v1.Y() + v1.Z()*v1.Z()
}

// SqrLength returns the squared length of v1
func (v1 *Vec3) SqrLength() float64 {
	return v1.X()*v1.X() + v1.Y()*v1.Y() + v1.Z()*v1.Z()
}

// all the operation are provided in two fashions:
// 		1. Vec3 class method taking one Vec3 ptr
// 		2. normal function taking two Vec3 ptr

// using variadics to support multi-variable Add, Sub

// Add use first argument as pivot vector, iterate to add rest of vectors
func Add(v1 *Vec3, vs ...*Vec3) *Vec3 {
	e0, e1, e2 := v1.X(), v1.Y(), v1.Z()
	for _, v := range vs {
		e0 += v.X()
		e1 += v.Y()
		e2 += v.Z()
	}
	sum := NewVec3(e0, e1, e2)
	return sum
}

// Add use first argument as pivot vector, iterate to add rest of vectors
func (v1 *Vec3) Add(vs ...*Vec3) *Vec3 {
	e0, e1, e2 := v1.X(), v1.Y(), v1.Z()
	for _, v := range vs {
		e0 += v.X()
		e1 += v.Y()
		e2 += v.Z()
	}
	sum := NewVec3(e0, e1, e2)
	return sum
}

// Sub use first argument as pivot vector, iterate to substract rest of vectors
func Sub(v1 *Vec3, vs ...*Vec3) *Vec3 {
	e0, e1, e2 := v1.X(), v1.Y(), v1.Z()
	for _, v := range vs {
		e0 -= v.X()
		e1 -= v.Y()
		e2 -= v.Z()
	}
	res := NewVec3(e0, e1, e2)
	return res
}

// Sub use first argument as pivot vector, iterate to substract rest of vectors
func (v1 *Vec3) Sub(vs ...*Vec3) *Vec3 {
	e0, e1, e2 := v1.X(), v1.Y(), v1.Z()
	for _, v := range vs {
		e0 -= v.X()
		e1 -= v.Y()
		e2 -= v.Z()
	}
	res := NewVec3(e0, e1, e2)
	return res
}

// Mut performs scalar multiplication on v1
func Mut(v1 *Vec3, s float64) *Vec3 {
	return NewVec3(
		v1.X()*s,
		v1.Y()*s,
		v1.Z()*s,
	)
}

// Mut performs scalar multiplication on v1
func (v1 *Vec3) Mut(s float64) *Vec3 {
	return NewVec3(
		v1.X()*s,
		v1.Y()*s,
		v1.Z()*s,
	)
}

// Div performs scalar division on v1
func Div(v1 *Vec3, s float64) *Vec3 {
	return NewVec3(
		v1.X()/s,
		v1.Y()/s,
		v1.Z()/s,
	)
}

// Div performs scalar division on v1
func (v1 *Vec3) Div(s float64) *Vec3 {
	return NewVec3(
		v1.X()/s,
		v1.Y()/s,
		v1.Z()/s,
	)
}

// Dot performs dot product, v1 as the pivot vector
func Dot(v1, v2 *Vec3) float64 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

// Dot performs dot product, v1 as the pivot vector
func (v1 *Vec3) Dot(v2 *Vec3) float64 {
	return v1.X()*v2.X() + v1.Y()*v2.Y() + v1.Z()*v2.Z()
}

// Crx performs the cross product, v1 as the pivot vector
func Crx(v1, v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.Y()*v2.Z()-v1.Z()*v2.Y(),
		v1.Z()*v2.X()-v1.X()*v2.Z(),
		v1.X()*v2.Y()-v1.Z()*v2.X(),
	)
}

// Crx performs the cross product, v1 as the pivot vector
func (v1 *Vec3) Crx(v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.Y()*v2.Z()-v1.Z()*v2.Y(),
		v1.Z()*v2.X()-v1.X()*v2.Z(),
		v1.X()*v2.Y()-v1.Z()*v2.X(),
	)
}

// UnitVector return the unit-length vector along v1 direction
func UnitVector(v1 *Vec3) *Vec3 {
	return v1.Div(v1.Length())
}

// UnitVector return the unit-length vector along v1 direction
func (v1 *Vec3) UnitVector() *Vec3 {
	return v1.Div(v1.Length())
}
