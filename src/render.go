package main

import (
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
	// fmt.Println(render.RandomScene("./scene.csv"))

	// // CPU profiling by default
	// defer profile.Start(profile.CPUProfile).Stop()

	nThread, csvPath, output := render.ArgParse()

	w := render.SceneParser(csvPath)

	pos := &vec3.Vec3{7, 7, 7}
	lookAt := &vec3.Vec3{1, 0.2, 1}
	up := &vec3.Vec3{0, 1, 0}

	sampler := render.NewSampler(nx, ny, finess, maxDepth, tMin, seed)
	if nThread != 1 {
		sampler.SetParallel(nThread)
	}
	sampler.SetCamera(fov, aspect, aperture, pos, lookAt, up)

	sampler.SetWorldObj(w)

	sampler.Render()

	sampler.Save(output)
}
