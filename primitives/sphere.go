package primitives

import (
	"math"
	"ray"
	"vector"
)

// Sphere has one centroid, and a radius
type Sphere struct {
	Center *vector.Vec3
	Radius float64
}

// NewSphere creates new Sphere obj
func NewSphere(x, y, z, radius float64) *Sphere {
	return &Sphere{
		Center: vector.NewVec3(x, y, z),
		Radius: radius,
	}
}

// Hit a sphere could result in two hit spots, whichever first should win
func (s *Sphere) Hit(r *ray.Ray, tMin, tMax float64) *Hit {
	oc := r.Origin.Sub(s.Center)
	a := r.Direct.SqrLength()
	b := vector.Dot(oc, r.Direct)
	c := oc.SqrLength() - s.Radius*s.Radius
	delta := b*b - a*c

	hit := &Hit{}

	if delta > 0 {

		if temp := (-b - math.Sqrt(delta)) / a; temp < tMax && temp > tMin {
			hit.T = temp
			hit.Point = r.PointAtScale(temp)
			hit.Normal = hit.Point.Sub(s.Center).Div(s.Radius)
			return hit
		}
		if temp := (-b + math.Sqrt(delta)) / a; temp < tMax && temp > tMin {
			hit.T = temp
			hit.Point = r.PointAtScale(temp)
			hit.Normal = hit.Point.Sub(s.Center).Div(s.Radius)
			return hit
		}
	}
	return nil
}
