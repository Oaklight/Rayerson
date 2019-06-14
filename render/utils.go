package render

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	pm "primitives"
	"ray"
	"runtime"
	"strconv"
	"strings"
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

func ArgParse() (nThread int, csvFile string) {
	helpMsg := `Usage: render [-p=[num of threads]] <csv file>
	<csv file> = The file that contains the batch of images that need to be processed.
	-p=[num of threads] = An optional flag to run the editor in its parallel version.
	You also have the option of specifying the number of threads
	[num of threads] = the number of workers in the program (including the main thread)
	`
	nThread = 1

	// widthPtr := flag.Bool("fork", false, "a bool")

	switch len(os.Args) {
	case 3:
		flag := strings.Split(os.Args[1], "=")
		if len(flag) == 2 {
			nThread, _ = strconv.Atoi(flag[1])
		} else {
			nThread = runtime.NumCPU()
		}
		// in case there is malicious input
		if nThread > 0 {
			// TODO: create threads for rendering
		}
		fallthrough
	case 2:
		csvFile = os.Args[len(os.Args)-1]
	default:
		fmt.Println(helpMsg)
		os.Exit(1)
	}

	return
}

func csvReadline(csvReader *bufio.Scanner) pm.Hitable {
	// currently only parse sphere
	if csvReader.Scan() {
		args := strings.Split(csvReader.Text(), ",")
		fmt.Println(args)
		x, _ := strconv.ParseFloat(args[0], 64)
		y, _ := strconv.ParseFloat(args[1], 64)
		z, _ := strconv.ParseFloat(args[2], 64)
		radius, _ := strconv.ParseFloat(args[3], 64)

		var material pm.Materials
		switch args[4] {
		case "Diffuse":
			r, _ := strconv.ParseFloat(args[5], 64)
			g, _ := strconv.ParseFloat(args[6], 64)
			b, _ := strconv.ParseFloat(args[7], 64)
			material = pm.NewDiffuse(&ray.Color{r, g, b})
		case "Metallic":
			r, _ := strconv.ParseFloat(args[5], 64)
			g, _ := strconv.ParseFloat(args[6], 64)
			b, _ := strconv.ParseFloat(args[7], 64)
			fuzz, _ := strconv.ParseFloat(args[8], 64)
			material = pm.NewMetallic(&ray.Color{r, g, b}, fuzz)
		case "Dielectric":
			idx, _ := strconv.ParseFloat(args[5], 64)
			material = pm.NewDielectric(idx)
		}
		return pm.NewSphere(x, y, z, radius, material)
	}
	return nil
}

// type Context struct {
// 	nThread int
// 	done    chan bool
// 	lock    *sync.Mutex
// 	turn    *sync.Cond
// 	fetched *sync.Cond
// 	wg      *sync.WaitGroup
// 	job     []*description
// 	jt      jobType
// }

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
