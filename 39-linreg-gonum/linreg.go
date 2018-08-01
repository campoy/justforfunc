package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func linearRegression(x, y []float64, alpha float64) (m, c float64) {
	var cost, dm, dc float64
	buf := make([]float64, len(x))

	for i := 0; i < iterations; i++ {
		cost, dm, dc = computeGradient(buf, x, y, m, c)
		m += -dm * alpha
		c += -dc * alpha
		if i%(iterations/100) == 0 {
			fmt.Printf("% 3.0f%% : cost(%.2f, %.2f) = %.2f\n",
				100*float64(i)/float64(iterations), m, c, cost)
		}
	}

	fmt.Printf("100%% : cost(%.2f, %.2f) = %.2f\n", m, c, cost)

	return m, c
}

func computeGradient(buf, x, y []float64, m, c float64) (cost, dm, dc float64) {
	copy(buf, x)
	floats.Scale(m, buf)
	floats.AddConst(c, buf)
	floats.Sub(buf, y)

	n := float64(len(x))
	cost = floats.Dot(buf, buf) / n
	dm = 2 * floats.Dot(x, buf) / n
	dc = 2 * floats.Sum(buf) / n
	return cost, dm, dc
}

func computeGradientLoop(buf, x, y []float64, m, c float64) (cost, dm, dc float64) {
	for i := range x {
		d := y[i] - (x[i]*m + c)
		cost += d * d
		dm += -x[i] * d
		dc += -d
	}
	n := float64(len(x))
	cost = cost / n
	dm = 2 * dm / n
	dc = 2 * dc / n
	return cost, dm, dc
}
