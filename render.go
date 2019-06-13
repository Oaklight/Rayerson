package main

import (
	"fmt"
	"os"
	pm "primitives"
	"ray"
	"sampler"
)

const (
	maxDepth = 50
	tMin     = 0.001
	seed     = 42
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func main() {

	nx, ny := 800, 400
	finess := 100

	sampler := sampler.NewSampler(nx, ny, finess, maxDepth, tMin, seed)
	sampler.SetCamera()

	// f, err := os.Create("out.ppm")
	// defer f.Close()
	// check(err, "Error opening file: %v\n")
	// // http://netpbm.sourceforge.net/doc/ppm.html
	// _, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)
	// check(err, "Error writting to file: %v\n")

	w := &pm.World{}
	w.Add(
		pm.NewSphere(0, 0, -1, 0.5, pm.NewLambertian(ray.NewColor(0.8, 0.3, 0.3))),
		pm.NewSphere(0, -100.5, -1, 100, pm.NewLambertian(ray.NewColor(0.8, 0.8, 0))),
		pm.NewSphere(1, 0, -1, 0.5, pm.NewMetallic(ray.NewColor(0.8, 0.6, 0.2), 0.3)),
		pm.NewSphere(-1, 0, -1, 0.5, pm.NewMetallic(ray.NewColor(0.8, 0.8, 0.8), 0.7)),
	)

	sampler.SetWorldObj(w)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			sampler.SamplePixel(i, j)
			// _, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)

			// check(err, "Error writting to file: %v\n")
		}
	}

	// SaveImg(imgOut, "out.png")
	sampler.Save("out.png")
}
