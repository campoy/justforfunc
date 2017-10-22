// Copyright 2017 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

// Run these benchmarks with
//   go test -bench=.

package main

import (
	"fmt"
	"testing"
)

var sizes = []int{64, 128, 256, 512}

func BenchmarkSeq(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				createSeq(size, size)
			}
		})
	}
}

func BenchmarkPixel(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				createPixel(size, size)
			}
		})
	}
}

func BenchmarkRow(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				createRow(size, size)
			}
		})
	}
}

func BenchmarkWorkers(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				createWorkers(size, size, false)
			}
		})
	}
}

func BenchmarkWorkersBuffered(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				createWorkers(size, size, true)
			}
		})
	}
}

func BenchmarkRowWorkers(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				createRowWorkers(size, size, false)
			}
		})
	}
}

func BenchmarkRowWorkersBuffered(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				createRowWorkers(size, size, true)
			}
		})
	}
}
