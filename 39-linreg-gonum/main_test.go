package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestComputeGradient(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{1, 1, 1, 1, 1}
	buf := make([]float64, len(x))
	m, c := 2.0, 3.0
	cost, dm, dc := computeGradient(buf, x, y, m, c)
	wcost, wm, wc := computeGradientLoop(buf, x, y, m, c)
	if cost != wcost {
		t.Errorf("expected cost %g; got %g", wcost, cost)
	}
	if dm != wm {
		t.Errorf("expected dm %g; got %g", wm, dm)
	}
	if dc != wc {
		t.Errorf("expected dc %g; got %g", wc, dc)
	}
}

func BenchmarkComputeGradient(b *testing.B) {
	x := make([]float64, 1000)
	y := make([]float64, 1000)
	for i := range x {
		x[i] = rand.Float64()
		y[i] = rand.Float64()
	}
	buf := make([]float64, len(x))

	for i := 0; i < 10; i++ {
		b.Run(fmt.Sprint(len(x)), func(b *testing.B) {
			fs := []struct {
				name string
				f    func(buf, x, y []float64, m, c float64) (cost, dm, dc float64)
			}{
				{"floats", computeGradient},
				{"loop", computeGradientLoop},
			}
			for _, f := range fs {
				b.Run(f.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						f.f(buf, x, y, 1, 2)
					}
				})
			}
		})

		x = append(x, x...)
		y = append(y, y...)
		buf = make([]float64, len(x))
	}
}
