package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func linearRegression(x, y []float64, alpha float64) (m, c float64) {
	for i := 0; i < iterations; i++ {
		dm, dc := computeGradient(x, y, m, c)
		m += -dm * alpha
		c += -dc * alpha
		fmt.Printf("cost(%.2f, %.2f) = %.2f\n", m, c, computeCost(x, y, m, c))
	}

	fmt.Printf("cost(%.2f, %.2f) = %.2f\n", m, c, computeCost(x, y, m, c))

	return m, c
}

func clone(x []float64) []float64 {
	c := make([]float64, len(x))
	copy(c, x)
	return c
}

func computeCost(x, y []float64, m, c float64) float64 {
	x, y = clone(x), clone(y)
	floats.Scale(m, x)    // m * x
	floats.AddConst(c, x) // m * x + c
	floats.Sub(y, x)      // y - (m*x + c)
	return floats.Dot(y, y) / float64(len(x))
}

func computeGradient(x, y []float64, m, c float64) (dm, dc float64) {
	origX := x
	x, y = clone(x), clone(y)
	floats.Scale(m, x)    // m * x
	floats.AddConst(c, x) // m * x + c
	floats.Sub(y, x)      // y - (m*x + c)

	f := -2 / float64(len(x))
	// cost/dm = -2/N * sum(x * (y - (m*x+c)))
	dm = f * floats.Dot(origX, y)
	// cost/dc = -2/N * sum((y - (m*x+c)))
	dc = f * floats.Sum(y)
	return dm, dc
}
