package main

import (
	"fmt"
	"testing"
)

func TestComputeGradient(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{1, 1, 1, 1, 1}
	m, c := 2.0, 3.0
	dm, dc := computeGradient(x, y, m, c)
	wm, wc := computeGradientLoop(x, y, m, c)
	if dm != wm {
		t.Errorf("expected dm %g; got %g", wm, dm)
	}
	if dc != wc {
		t.Errorf("expected dc %g; got %g", wc, dc)
	}
}

func computeGradientLoop(x, y []float64, m, c float64) (dm, dc float64) {
	// cost = 1/N * sum((y - (m*x+c))^2)
	// cost/dm = 2/N * sum(-x * (y - (m*x+c)))
	// cost/dc = 2/N * sum(-(y - (m*x+c)))
	fmt.Println("loop")
	fmt.Println("x", x)
	for i := range x {
		fmt.Println("x[i]", x[i])
		fmt.Println("x[i]*m", x[i]*m)
		d := y[i] - (x[i]*m + c)
		fmt.Println("d", d)
		dm += -x[i] * d
		fmt.Println("dm", x[i]*d)
		dc += -d
	}
	n := float64(len(x))
	return 2 / n * dm, 2 / n * dc
}
