package render

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	pm "primitives"
	"ray"
	"time"
	vec3 "vector"
)

// Sampler performs the color sampling on pixel level
type Sampler struct {
	width, height    int
	finess, maxDepth int
	isParallel       bool
	tMin, tMax       float64
	imgOut           *image.RGBA64
	rnd              *rand.Rand
	cam              *ray.Camera
	world            *pm.World
}

// NewSampler creates a new sampler for rendering
func NewSampler(width, height, finess, maxDepth int, tMin float64, seed ...int) *Sampler {
	s := Sampler{
		width:    width,
		height:   height,
		finess:   finess,
		maxDepth: maxDepth,
		tMin:     tMin,
		tMax:     math.MaxFloat64,
		imgOut:   image.NewRGBA64(image.Rect(0, 0, width, height)),
	}
	var src rand.Source
	switch len(seed) {
	case 0:
		src = rand.NewSource(time.Now().UnixNano())
	default:
		src = rand.NewSource(int64(seed[0]))
	}
	s.rnd = rand.New(src)
	return &s
}

// SetCamera customize the camera model with given parameters
// ** lookAt is a point
func (s *Sampler) SetCamera(fov, aspect, aperture float64, pos, lookAt, up *vec3.Vec3) {
	s.cam = ray.NewCamera(fov, aspect, aperture, *pos, *lookAt, *up)
}

// SetWorldObj sets up the world of hitable objects
func (s *Sampler) SetWorldObj(world *pm.World) {
	s.world = world
}

// Save saves the image to the given file
func (s *Sampler) Save(filePath string) error {

	outWriter, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outWriter.Close()

	err = png.Encode(outWriter, s.imgOut)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sampler) color4Ray(r *ray.Ray, depth int) *ray.Color {
	if hit := s.world.Hit(r, s.tMin, s.tMax); hit != nil {

		if bounced := hit.Materials.Bounce(r, hit, s.rnd); bounced != nil && depth < s.maxDepth {
			newColor := s.color4Ray(bounced, depth+1)
			return hit.Color().Mul(newColor)
		}
		return &ray.Opaque
	}

	unitDirect := r.Direct.Normalize()
	t := 0.5 * (unitDirect.Y + 1)
	return ray.Transparent.MulScalar(1.0 - t).Add(ray.Opaque.MulScalar(t))
}

// SamplePixel yields the color for given coordinate (x, y)
func (s *Sampler) SamplePixel(x, y int) color.RGBA64 {
	col := &ray.Color{}

	// anti-aliasing
	// refine the color by sampling around each pixel, up to given finess
	for rf := 0; rf < s.finess; rf++ {
		u := (float64(x) + rand.Float64()) / float64(s.width)
		v := (float64(y) + rand.Float64()) / float64(s.height)
		r := s.cam.GetRay(u, v, s.rnd)
		col = col.Add(s.color4Ray(r, 0))
	}
	col = col.DivScalar(float64(s.finess))
	rgba64 := col.RGBA64()
	s.imgOut.SetRGBA64(x, s.height-y, rgba64)

	return rgba64
}
