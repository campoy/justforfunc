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
)

var iterations int

func main() {
	flag.IntVar(&iterations, "n", 1000, "number of iterations")
	flag.Parse()

	x, y, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}

	err = plotData("out.png", x, y)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

type xy struct{ x, y float64 }

func readData(path string) (x, y []float64, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var a, b float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &a, &b)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		x = append(x, a)
		y = append(y, b)
	}
	if err := s.Err(); err != nil {
		return nil, nil, fmt.Errorf("could not scan: %v", err)
	}
	return x, y, nil
}

type zip struct{ x, y []float64 }

func (z zip) Len() int                { return len(z.x) }
func (z zip) XY(i int) (x, y float64) { return z.x[i], z.y[i] }

func plotData(path string, x, y []float64) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}

	xys := zip{x, y}
	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	m, c := linearRegression(x, y, 0.01)

	// create fake linear regression result
	l, err := plotter.NewLine(plotter.XYs{
		{3, 3*m + c}, {20, 20*m + c},
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
