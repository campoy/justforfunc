package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"

	"github.com/campoy/justforfunc/38-linreg-vanilla/linreg"
)

func main() {
	iterations := flag.Int("n", 1000, "number of iterations")
	flag.Parse()

	xs, ys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}

	err = plotData("out.png", xs, ys, *iterations)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

func readData(path string) (xs, ys []float64, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xs = append(xs, x)
		ys = append(ys, y)
	}
	if err := s.Err(); err != nil {
		return nil, nil, fmt.Errorf("could not scan: %v", err)
	}
	return xs, ys, nil
}

type xyer struct{ xs, ys []float64 }

func (x xyer) Len() int                    { return len(x.xs) }
func (x xyer) XY(i int) (float64, float64) { return x.xs[i], x.ys[i] }

func plotData(path string, xs, ys []float64, iterations int) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}

	// create scatter with all data points
	s, err := plotter.NewScatter(xyer{xs, ys})
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	x, c := linreg.LinearRegression(xs, ys, iterations, 0.01)

	// create fake linear regression result
	l, err := plotter.NewLine(plotter.XYs{
		{3, 3*x + c}, {20, 20*x + c},
	})
	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}
