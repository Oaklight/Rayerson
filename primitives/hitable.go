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
	Materials
}

// Hitable requires all hitable objects to have a Hit function
type Hitable interface {
	Hit(r *ray.Ray, tMin, tMax float64) *Hit
}

// World defines a series of Hitable objects
type World []Hitable

// func NewWorld(size int) *World {
// 	return make([]Hitable, 0, size).(World)
// }

// Add appends the given variable(s) to a world list
func (w *World) Add(hs ...Hitable) {
	for _, each := range hs {
		*w = append(*w, each)
	}
}

// Count returns how many elements are there in the world
func (w *World) Count() int {
	return len(*w)
}

// Hit iterates over the world objects and try to hit each one.
func (w *World) Hit(r *ray.Ray, tMin, tMax float64) *Hit {
	closet := tMax
	var record *Hit

	for _, each := range *w {
		if each != nil {
			// if some node already intersected with a much nearer object,
			// closet will be updated and that would block anything farther
			// to be hit successfully
			if hit := each.Hit(r, tMin, closet); hit != nil {
				closet = hit.T
				record = hit
			}
		}
	}
	return record
}
