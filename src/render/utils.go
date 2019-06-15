package render

import (
	"bufio"
	"fmt"
	"os"
	pm "primitives"
	"ray"
	"runtime"
	"strconv"
	"strings"
)

func ArgParse() (nThread int, csvFile string, outFile string) {
	helpMsg := `Usage: render [-p=[num of threads]] <csv file> <output file>
	<csv file> = The file that contains the batch of images that need to be processed.
	-p=[num of threads] = An optional flag to run the editor in its parallel version.
	You also have the option of specifying the number of threads
	[num of threads] = the number of workers in the program (including the main thread)
	`
	nThread = 1

	// widthPtr := flag.Bool("fork", false, "a bool")

	switch len(os.Args) {
	case 4:
		flag := strings.Split(os.Args[1], "=")
		if len(flag) == 2 {
			nThread, _ = strconv.Atoi(flag[1])
		} else {
			nThread = runtime.NumCPU()
		}
		// in case there is malicious input
		if nThread > 0 {
			// TODO: create threads for rendering
			// initWorkerPool(nThread)
		}
		fallthrough
	case 3:
		csvFile = os.Args[len(os.Args)-2]
		outFile = os.Args[len(os.Args)-1]
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

// func initWorkerPool(nThread int) {

// }
