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

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	output = "out.png"
	width  = 1024
	height = 1024
)

func main() {
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	img := create(width, height)

	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func create(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			m.Set(i, j, pixel(i, j, width, height))
		}
	}
	return m
}

func pixel(i, j, width, height int) color.Color {
	xi := norm(i, height, -2.5, 1)
	yi := norm(j, width, -1, 1)

	const maxI = 1000
	x, y := 0., 0.
	for i := 0; (x*x+y*y < 4) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
