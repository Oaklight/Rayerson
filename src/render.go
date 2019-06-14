package main

import (
	"fmt"
	"render"
	vec3 "vector"
)

const (
	// ray tracing
	maxDepth = 50
	tMin     = 0.001
	finess   = 100
	seed     = 42
	// camera
	nx, ny = 800, 400
	aspect = nx / ny
	// larger for wider view
	fov = 40
	// larger for better defocus
	aperture = 0.1
)

func main() {
	// // CPU profiling by default
	// defer profile.Start(profile.CPUProfile).Stop()

	_, csvPath := render.ArgParse()
	// nThread, csvPath := argParse()
	w := render.SceneParser(csvPath)
	// fmt.Println(w)
	for _, each := range *w {
		fmt.Println(each)
	}

	pos := &vec3.Vec3{3, 3, 2}
	lookAt := &vec3.Vec3{0, 0, 1}
	up := &vec3.Vec3{0, 1, 0}

	sampler := render.NewSampler(nx, ny, finess, maxDepth, tMin, seed)
	sampler.SetCamera(fov, aspect, aperture, pos, lookAt, up)

	// w := &pm.World{}
	// w.Add(
	// 	pm.NewSphere(0, 0, -1, 0.5, pm.NewDiffuse(ray.NewColor(0.1, 0.2, 0.5))),
	// 	pm.NewSphere(0, -100.5, -1, 100, pm.NewDiffuse(ray.NewColor(0.8, 0.8, 0))),
	// 	pm.NewSphere(1, 0, -1, 0.5, pm.NewMetallic(ray.NewColor(0.8, 0.6, 0.2), 0.3)),
	// 	// pm.NewSphere(-1, 0, -1, 0.5, pm.NewMetallic(ray.NewColor(0.3, 0.8, 0.5), 0.8)),
	// 	pm.NewSphere(-1, 0, -1, 0.5, pm.NewDielectric(1.5)),
	// 	pm.NewSphere(-1, 0, -1, -0.45, pm.NewDielectric(1.5)),
	// )

	sampler.SetWorldObj(w)

	for j := 1; j < ny; j++ {
		for i := 0; i < nx; i++ {
			sampler.SamplePixel(i, j)
		}
	}

	sampler.Save("out.png")
}
