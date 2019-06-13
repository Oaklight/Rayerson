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

// ========================= DiffuseMaterial =========================

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

// ========================= MetallicMaterial =========================

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
	reflected := r.Direct.Reflect(hit.Normal)
	if reflected.Dot(hit.Normal) > 0 {
		fuzzed := reflected.Add(vec3.RandUnitVec3(rnd).MulScalar(m.Fuzz))
		return ray.NewRay(hit.Point, fuzzed)
	}
	return nil
}

// ========================= DielectricMaterial =========================

// DielectricMaterial is transparent, with varying refractance index
type DielectricMaterial struct {
	refIdx      float64
	attenuation *ray.Color
}

func NewDielectric(refIdx float64) *DielectricMaterial {
	return &DielectricMaterial{
		refIdx:      refIdx,
		attenuation: ray.NewColor(1, 1, 1),
	}
}

func (d *DielectricMaterial) Color() *ray.Color {
	return d.attenuation
}

// Schlick's approximation: https://en.wikipedia.org/wiki/Schlick%27s_approximation
func (d *DielectricMaterial) schlick(cosine float64) float64 {
	r0 := (1.0 - d.refIdx) / (1.0 + d.refIdx)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow((1.0-cosine), 5)
}

func (d *DielectricMaterial) Bounce(r *ray.Ray, hit *Hit, rnd *rand.Rand) *ray.Ray {
	var ratio, cosine, reflectChance float64
	var normalOutward *vec3.Vec3

	divisor := r.Direct.Length()
	cosine = r.Direct.Dot(hit.Normal) * d.refIdx / divisor
	if cosine > 0 {
		normalOutward = hit.Normal.Negate()
		ratio = d.refIdx
	} else {
		normalOutward = hit.Normal
		ratio = 1.0 / d.refIdx
		cosine = -cosine
	}

	refracted := r.Direct.Refract(normalOutward, ratio)

	if refracted != nil {
		reflectChance = d.schlick(cosine)
	} else {
		reflectChance = 1.0
	}
	if rnd.Float64() < reflectChance {
		reflected := r.Direct.Reflect(hit.Normal)
		return ray.NewRay(hit.Point, reflected)
	}

	return ray.NewRay(hit.Point, refracted)
}
