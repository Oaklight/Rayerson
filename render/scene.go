package render

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	pm "primitives"
	"ray"
	"time"
	vec3 "vector"
)

func SceneParser(csvPath string) *pm.World {
	csvFile, _ := os.Open(csvPath)
	defer csvFile.Close()
	scanner := bufio.NewScanner(csvFile)
	worldObj := pm.World{}
	for {

		if obj := csvReadline(scanner); obj != nil {

			worldObj.Add(obj)

		} else {
			break
		}
	}
	return &worldObj
}

// RandomScene returns a 'random' scene
func RandomScene(csvPath string) *pm.World {
	rand.Seed(time.Now().UnixNano())

	csvFile, err := os.Create(csvPath)
	if err != nil {
		panic("File creation failed")
	}
	defer csvFile.Close()

	world := pm.World{}
	floor := pm.NewSphere(0, -1000, -1, 1000, pm.NewDiffuse(ray.NewColor(0.5, 0.5, 0.5)))
	world.Add(floor)
	radius := 0.2
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			material := rand.Float64()

			center := vec3.Vec3{
				X: float64(a) + 0.9*rand.Float64(),
				Y: radius,
				Z: float64(b) + 0.9*rand.Float64(),
			}

			if center.Sub(&vec3.Vec3{4, radius, 0}).Length() > 0.9 {
				if material < 0.8 {
					r, g, b := rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64()
					diffuse := pm.NewSphere(center.X, center.Y, center.Z, radius,
						pm.NewDiffuse(&ray.Color{r, g, b}))

					world.Add(diffuse)
					fmt.Fprintf(csvFile,
						"%f,%f,%f,%f,%s,%f,%f,%f\n",
						center.X, center.Y, center.Z, radius, "Diffuse", r, g, b)

				} else if material < 0.95 {
					r, g, b := 0.5*(1.0+rand.Float64()), 0.5*(1.0+rand.Float64()), 0.5*(1.0+rand.Float64())
					fuzz := 0.5 + rand.Float64()
					metal := pm.NewSphere(center.X, center.Y, center.Z, radius,
						pm.NewMetallic(&ray.Color{r, g, b}, fuzz))

					world.Add(metal)
					fmt.Fprintf(csvFile,
						"%f,%f,%f,%f,%s,%f,%f,%f,%f\n",
						center.X, center.Y, center.Z, radius, "Metallic", r, g, b, fuzz)

				} else {
					idx := 1.5
					glass := pm.NewSphere(center.X, center.Y, center.Z, radius,
						pm.NewDielectric(idx))

					world.Add(glass)
					fmt.Fprintf(csvFile,
						"%f,%f,%f,%f,%s,%f\n",
						center.X, center.Y, center.Z, radius, "Dielectric", idx)

				}
			}
		}
	}

	glass := pm.NewSphere(0, 1, 0, 1.0, pm.NewDielectric(1.5))
	diffuse := pm.NewSphere(-4, 1, 0, 1.0, pm.NewDiffuse(&ray.Color{0.4, 0, 0.1}))
	metal := pm.NewSphere(4, 1, 0, 1.0, pm.NewMetallic(&ray.Color{0.7, 0.6, 0.5}, 0))

	world.Add(glass, diffuse, metal)
	fmt.Fprintf(csvFile,
		"%f,%f,%f,%f,%s,%f\n",
		0.0, 1.0, 0.0, 1.0, "Dielectric", 1.5)
	fmt.Fprintf(csvFile,
		"%f,%f,%f,%f,%s,%f,%f,%f\n",
		-4.0, 1.0, 0.0, 1.0, "Diffuse", 0.4, 0.0, 0.1)
	fmt.Fprintf(csvFile,
		"%f,%f,%f,%f,%s,%f,%f,%f,%f\n",
		4.0, 1.0, 0.0, 1.0, "Metallic", 0.7, 0.6, 0.5, 0.0)
	return &world
}
