package primitives

import (
	"math"
	"math/rand"
	"ray"
	vec3 "vector"
)

// Materials defines the interface type of different materials
type Materials interface {
	Bounce(r *ray.Ray, hit *Hit, rnd *rand.Rand) *ray.Ray
	Color() *ray.Color
}

// ============================ DiffuseMaterial material ============================

// DiffuseMaterial type
type DiffuseMaterial struct {
	Albedo *ray.Color
}

func NewLambertian(color *ray.Color) *DiffuseMaterial {
	return &DiffuseMaterial{color}
}

func (l *DiffuseMaterial) Color() *ray.Color {
	return l.Albedo
}

func (l *DiffuseMaterial) Bounce(r *ray.Ray, hit *Hit, rnd *rand.Rand) *ray.Ray {
	scattered := hit.Normal.Add(vec3.RandUnitVec3(rnd))
	return ray.NewRay(hit.Point, scattered)
}

// ============================ MetallicMaterial material ============================

// MetallicMaterial type
type MetallicMaterial struct {
	Albedo *ray.Color
	Fuzz   float64
}

func NewMetallic(color *ray.Color, fuzziness float64) *MetallicMaterial {
	return &MetallicMaterial{
		Albedo: color,
		Fuzz:   math.Min(fuzziness, 1),
	}
}

func (m *MetallicMaterial) Color() *ray.Color {
	return m.Albedo
}

func (m *MetallicMaterial) Bounce(r *ray.Ray, hit *Hit, rnd *rand.Rand) *ray.Ray {
	reflected := r.Direct.Normalize().Reflect(hit.Normal)
	if reflected.Dot(hit.Normal) > 0 {
		fuzzed := reflected.Add(vec3.RandUnitVec3(rnd).MulScalar(m.Fuzz))
		return ray.NewRay(hit.Point, fuzzed)
	}
	return nil
}
