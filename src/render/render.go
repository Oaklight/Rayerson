package render

import (
	"fmt"
	"image/color"
	"sync"
)

const nThread = 2
const bufferSize = 2

type Pixel struct {
	X, Y  int
	Color *color.RGBA64
}

func (s *Sampler) Render() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	if s.isParallel {
		done := make(chan interface{})
		defer close(done)

		target := s.width * s.height

		blankPixelStream := make(chan Pixel, 10)
		defer close(blankPixelStream)

		workers := make([]<-chan int, nThread)
		for i := range workers {
			workers[i] = s.worker(i, done, blankPixelStream)
		}

		progress := collector(target, done, workers...)

		go func() {
			for j := 0; j < s.height; j++ {
				for i := 0; i < s.width; i++ {
					blankPixelStream <- Pixel{i, j, nil}
				}
			}
		}()
		completed := 0
		for p := range progress {
			// s.ImgOut.SetRGBA64(pixel.X, s.height-pixel.Y, *pixel.Color)
			if completed += p; completed == target {
				fmt.Println("all pixel rendered")
				break
			}
			//
		}
	} else {
		for j := 1; j < s.height; j++ {
			for i := 0; i < s.width; i++ {
				rgba64 := s.SamplePixel(i, j)
				s.ImgOut.SetRGBA64(i, s.height-j, rgba64)
			}
		}
	}
}

func (s *Sampler) worker(id int, done <-chan interface{}, blankPixelStream <-chan Pixel) <-chan int {
	progressStream := make(chan int, 10)
	go func() {
		defer close(progressStream)
		for p := range blankPixelStream {
			s.SamplePixel(p.X, p.Y)
			// p.Color = &c
			progressStream <- 1

			select {
			case <-done:
				return
			default:
			}
		}
	}()
	return progressStream
}

func collector(target int, done <-chan interface{}, progressStream ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedStream := make(chan int, 100)

	multiplex := func(px <-chan int) {
		defer wg.Done()
		for p := range px {
			select {
			case <-done:
				return
			case multiplexedStream <- p:
			}
		}
	}

	wg.Add(len(progressStream))
	for _, c := range progressStream {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
