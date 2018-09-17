package linreg_test

import (
	"math"
	"testing"

	"github.com/campoy/justforfunc/38-linreg-vanilla/linreg"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestGradient(t *testing.T) {
	tt := []struct {
		name         string
		x, y         []float64
		m, c         float64
		cost, dm, dc float64
	}{
		{
			"identify function with zero error",
			[]float64{0, 1, 2, 3, 4}, []float64{0, 1, 2, 3, 4},
			1.0, 0.0,
			0.0, 0.0, 0.0,
		},
		{
			"identify function with approximated with flat zero line",
			[]float64{0, 1, 2, 3, 4}, []float64{0, 1, 2, 3, 4},
			0.0, 0.0,
			6.0, -12.0, -4.0,
		},
		{
			"random points approximated with identity function",
			[]float64{0, 1, 2, 3, 4}, []float64{2.315, 0.3235, 2.212, 1.2235, 4.1234},
			1.0, 0.0,
			1.806600, 2.035360, -0.078960,
		},
	}

	opt := cmpopts.EquateApprox(0, 0.001)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cost, dm, dc := linreg.Gradient(tc.x, tc.y, tc.m, tc.c)
			if !cmp.Equal(tc.cost, cost, opt) {
				t.Fatalf("cost should be %.10f, got %.10f", tc.cost, cost)
			}
			if !cmp.Equal(tc.dm, dm, opt) {
				t.Fatalf("dm should be %.10f, got %.10f", tc.dm, dm)
			}
			if !cmp.Equal(tc.dc, dc, opt) {
				t.Fatalf("dc should be %.10f, got %.10f", tc.dc, dc)
			}
		})
	}
}

func TestLinearRegression(t *testing.T) {
	tt := []struct {
		name  string
		x     []float64
		f     func(float64) float64
		n     int
		alpha float64
		m, c  float64
	}{
		{
			"approximating identify function",
			[]float64{0, 1, 2, 3, 4}, func(x float64) float64 { return x },
			10000, 0.01,
			1.0, 0.0,
		},
		{
			"approximating square function",
			[]float64{0, 1, 2, 3, 4}, func(x float64) float64 { return x * x },
			10000, 0.01,
			4.0, -2.0,
		},
		{
			"approximating 2x+pi function",
			[]float64{0, 1, 2, 3, 4}, func(x float64) float64 { return 2*x + math.Pi },
			10000, 0.01,
			2.0, math.Pi,
		},
	}

	opt := cmpopts.EquateApprox(0, 0.001)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			y := make([]float64, len(tc.x))
			for i, v := range tc.x {
				y[i] = tc.f(v)
			}
			m, c := linreg.LinearRegression(tc.x, y, tc.n, tc.alpha)
			if !cmp.Equal(m, tc.m, opt) {
				t.Fatalf("bad value for m: %s", cmp.Diff(m, tc.m))
			}
			if !cmp.Equal(c, tc.c, opt) {
				t.Fatalf("bad value for c: %s", cmp.Diff(c, tc.c))
			}
		})
	}
}
