package vector

import "math"

type Vec3 struct {
	e [3]float64
}

func NewVec3(e0, e1, e2 float64) *Vec3 {
	var v1 Vec3
	v1.e[0] = e0
	v1.e[1] = e1
	v1.e[2] = e2
	return &v1
}

func (v1 *Vec3) X() float64 {
	return v1.e[0]
}

func (v1 *Vec3) Y() float64 {
	return v1.e[1]
}

func (v1 *Vec3) Z() float64 {
	return v1.e[2]
}

func (v1 *Vec3) R() float64 {
	return v1.e[0]
}

func (v1 *Vec3) G() float64 {
	return v1.e[1]
}

func (v1 *Vec3) B() float64 {
	return v1.e[2]
}

func (v1 *Vec3) At(i int) float64 {
	return v1.e[i]
}

// all the operation are provided in two fashions:
// 		1. Vec3 class method taking one Vec3 ptr
// 		2. normal function taking two Vec3 ptr

func Add(v1, v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.e[0]+v2.e[0],
		v1.e[1]+v2.e[1],
		v1.e[2]+v2.e[2],
	)
}

func (v1 *Vec3) Add(v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.e[0]+v2.e[0],
		v1.e[1]+v2.e[1],
		v1.e[2]+v2.e[2],
	)
}

func Sub(v1, v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.e[0]-v2.e[0],
		v1.e[1]-v2.e[1],
		v1.e[2]-v2.e[2],
	)
}

func (v1 *Vec3) Sub(v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.e[0]-v2.e[0],
		v1.e[1]-v2.e[1],
		v1.e[2]-v2.e[2],
	)
}

func Mut(v1 *Vec3, s float64) *Vec3 {
	return NewVec3(
		v1.e[0]*s,
		v1.e[1]*s,
		v1.e[2]*s,
	)
}

func (v1 *Vec3) Mut(s float64) *Vec3 {
	return NewVec3(
		v1.e[0]*s,
		v1.e[1]*s,
		v1.e[2]*s,
	)
}

func Div(v1 *Vec3, s float64) *Vec3 {
	return NewVec3(
		v1.e[0]/s,
		v1.e[1]/s,
		v1.e[2]/s,
	)
}

func (v1 *Vec3) Div(s float64) *Vec3 {
	return NewVec3(
		v1.e[0]/s,
		v1.e[1]/s,
		v1.e[2]/s,
	)
}

func Dot(v1, v2 *Vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}
func (v1 *Vec3) Dot(v2 *Vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

func Crx(v1, v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.e[1]*v2.e[2]-v1.e[2]*v2.e[1],
		v1.e[2]*v2.e[0]-v1.e[0]*v2.e[2],
		v1.e[0]*v2.e[1]-v1.e[2]*v2.e[0],
	)
}

func (v1 *Vec3) Crx(v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.e[1]*v2.e[2]-v1.e[2]*v2.e[1],
		v1.e[2]*v2.e[0]-v1.e[0]*v2.e[2],
		v1.e[0]*v2.e[1]-v1.e[2]*v2.e[0],
	)
}

func Zeros() *Vec3 {
	return NewVec3(0, 0, 0)
}

func Ones() *Vec3 {
	return NewVec3(1, 1, 1)
}

func Length(v1 *Vec3) float64 {
	return math.Sqrt(v1.e[0]*v1.e[0] + v1.e[1] + v1.e[1] + v1.e[2]*v1.e[2])
}

func (v1 *Vec3) Length() float64 {
	return math.Sqrt(v1.e[0]*v1.e[0] + v1.e[1] + v1.e[1] + v1.e[2]*v1.e[2])
}

func SqrLength(v1 *Vec3) float64 {
	return v1.e[0]*v1.e[0] + v1.e[1] + v1.e[1] + v1.e[2]*v1.e[2]
}

func (v1 *Vec3) SqrLength() float64 {
	return v1.e[0]*v1.e[0] + v1.e[1] + v1.e[1] + v1.e[2]*v1.e[2]
}

func UnitVector(v1 *Vec3) *Vec3 {
	return v1.Div(v1.Length())
}

func (v1 *Vec3) UnitVector() *Vec3 {
	return v1.Div(v1.Length())
}
