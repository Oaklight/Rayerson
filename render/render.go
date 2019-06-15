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

		blankPixelStream := make(chan Pixel)
		defer close(blankPixelStream)

		workers := make([]<-chan Pixel, nThread)
		for i := range workers {
			workers[i] = s.worker(i, done, blankPixelStream)
		}

		pixels := collector(target, done, workers...)

		go func() {
			for j := 0; j < s.height; j++ {
				for i := 0; i < s.width; i++ {
					blankPixelStream <- Pixel{i, j, nil}
				}
			}
		}()
		completed := 0
		for pixel := range pixels {
			// fmt.Println(integer)
			s.ImgOut.SetRGBA64(pixel.X, s.height-pixel.Y, *pixel.Color)
			// fmt.Println(completed, target)
			if completed++; completed == target {
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

func (s *Sampler) worker(id int, done <-chan interface{}, blankPixelStream <-chan Pixel) <-chan Pixel {
	pixelStream := make(chan Pixel)
	go func() {
		defer close(pixelStream)
		for p := range blankPixelStream {
			c := s.SamplePixel(p.X, p.Y)
			p.Color = &c
			pixelStream <- p

			select {
			case <-done:
				return
			default:
			}
		}
	}()
	return pixelStream
}

func collector(target int, done <-chan interface{}, pixelStreams ...<-chan Pixel) <-chan Pixel {
	var wg sync.WaitGroup
	multiplexedStream := make(chan Pixel)

	multiplex := func(px <-chan Pixel) {
		defer wg.Done()
		for i := range px {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(pixelStreams))
	for _, c := range pixelStreams {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
