package primitives

import (
	"math"
	"ray"
	vec3 "vector"
)

// Sphere has one centroid, and a radius
type Sphere struct {
	Center   *vec3.Vec3
	Radius   float64
	Material Materials
}

// NewSphere creates new Sphere obj
func NewSphere(x, y, z, radius float64, m Materials) *Sphere {
	return &Sphere{
		Center:   &vec3.Vec3{x, y, z},
		Radius:   radius,
		Material: m,
	}
}

// Hit a sphere could result in two hit spots, whichever first should win
func (s *Sphere) Hit(r *ray.Ray, tMin, tMax float64) *Hit {
	oc := r.Origin.Sub(s.Center)
	a := r.Direct.Dot(r.Direct)
	b := vec3.Dot(oc, r.Direct)
	c := oc.Dot(oc) - s.Radius*s.Radius
	delta := b*b - a*c

	hit := &Hit{Materials: s.Material}

	if delta > 0 {
		if temp := (-b - math.Sqrt(delta)) / a; temp < tMax && temp > tMin {
			hit.T = temp
			hit.Point = r.PointAtScale(temp)
			hit.Normal = hit.Point.Sub(s.Center).DivScalar(s.Radius)
			return hit
		}
		if temp := (-b + math.Sqrt(delta)) / a; temp < tMax && temp > tMin {
			hit.T = temp
			hit.Point = r.PointAtScale(temp)
			hit.Normal = hit.Point.Sub(s.Center).DivScalar(s.Radius)
			return hit
		}
	}
	return nil
}
