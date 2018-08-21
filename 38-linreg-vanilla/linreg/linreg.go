// Package linreg provides a basic implementation of linear regression
// with gradient descent on two dimensional data.
package linreg

import "fmt"

// LinearRegression runs the requested number of iterations of gradient
// descent and returns the latest approximated coefficients.
func LinearRegression(xs, ys []float64, iterations int, alpha float64) (m, c float64) {
	for i := 0; i < iterations; i++ {
		cost, dm, dc := Gradient(xs, ys, m, c)
		m += -dm * alpha
		c += -dc * alpha
		if (10 * i % iterations) == 0 {
			fmt.Printf("cost(%.2f, %.2f) = %.2f\n", m, c, cost)
		}
	}

	return m, c
}

// Gradient computes the cost function and its gradients.
func Gradient(xs, ys []float64, m, c float64) (cost, dm, dc float64) {
	for i := range xs {
		d := ys[i] - (xs[i]*m + c)
		cost += d * d
		dm += -xs[i] * d
		dc += -d
	}
	n := float64(len(xs))
	return cost / n, 2 / n * dm, 2 / n * dc
}
