// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This program generates Mandelbrot fractals using different
// concurrency patterns.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"sync"
)

const (
	output     = "out.png"
	width      = 2048
	height     = 2048
	numWorkers = 8
)

func main() {
	// uncomment these lines to generate traces into stdout.
	// trace.Start(os.Stdout)
	// defer trace.Stop()

	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	img := createCol(width, height)

	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

// createSeq fills one pixel at a time.
func createSeq(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			m.Set(i, j, pixel(i, j, width, height))
		}
	}
	return m
}

// createPixel creates one goroutine per pixel.
func createPixel(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var w sync.WaitGroup
	w.Add(width * height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			go func(i, j int) {
				m.Set(i, j, pixel(i, j, width, height))
				w.Done()
			}(i, j)
		}
	}
	w.Wait()
	return m
}

// createCol creates one goroutine per column.
func createCol(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var w sync.WaitGroup
	w.Add(width)
	for i := 0; i < width; i++ {
		go func(i int) {
			for j := 0; j < height; j++ {
				m.Set(i, j, pixel(i, j, width, height))
			}
			w.Done()
		}(i)
	}
	w.Wait()
	return m
}

// createWorkers creates numWorkers workers and uses a channel to pass each pixel.
func createWorkers(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	type px struct{ x, y int }
	c := make(chan px)

	var w sync.WaitGroup
	for n := 0; n < numWorkers; n++ {
		w.Add(1)
		go func() {
			for px := range c {
				m.Set(px.x, px.y, pixel(px.x, px.y, width, height))
			}
			w.Done()
		}()
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c <- px{i, j}
		}
	}
	close(c)
	w.Wait()
	return m
}

// createWorkersBuffered creates numWorkers workers and uses a buffered channel to pass each pixel.
func createWorkersBuffered(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	type px struct{ x, y int }
	c := make(chan px, width*height)

	var w sync.WaitGroup
	for n := 0; n < numWorkers; n++ {
		w.Add(1)
		go func() {
			for px := range c {
				m.Set(px.x, px.y, pixel(px.x, px.y, width, height))
			}
			w.Done()
		}()
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c <- px{i, j}
		}
	}
	close(c)
	w.Wait()
	return m
}

// createColWorkers creates numWorkers workers and uses a channel to pass each column.
func createColWorkers(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	c := make(chan int)

	var w sync.WaitGroup
	for n := 0; n < numWorkers; n++ {
		w.Add(1)
		go func() {
			for i := range c {
				for j := 0; j < height; j++ {
					m.Set(i, j, pixel(i, j, width, height))
				}
			}
			w.Done()
		}()
	}

	for i := 0; i < width; i++ {
		c <- i
	}

	close(c)
	w.Wait()
	return m
}

// createColWorkersBuffered creates numWorkers workers and uses a buffered channel to pass each column.
func createColWorkersBuffered(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	c := make(chan int, width)

	var w sync.WaitGroup
	for n := 0; n < numWorkers; n++ {
		w.Add(1)
		go func() {
			for i := range c {
				for j := 0; j < height; j++ {
					m.Set(i, j, pixel(i, j, width, height))
				}
			}
			w.Done()
		}()
	}

	for i := 0; i < width; i++ {
		c <- i
	}

	close(c)
	w.Wait()
	return m
}

// pixel returns the color of a Mandelbrot fractal at the given point.
func pixel(i, j, width, height int) color.Color {
	// Play with this constant to increase the complexity of the fractal.
	// In the justforfunc.com video this was set to 4.
	const complexity = 1024

	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, -1, 1)

	const maxI = 1000
	x, y := 0., 0.

	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
