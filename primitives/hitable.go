package primitives

import (
	"ray"
	"vector"
)

// Hit contains the scaling factor T along ray direction, and
// contains the intersection point, and the surface normal at that point.
type Hit struct {
	T             float64
	Point, Normal *vector.Vec3
}

// Hitable requires all hitable objects to have a Hit function
type Hitable interface {
	Hit(r *ray.Ray, tMin, tMax float64) *Hit
}
