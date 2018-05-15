package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
)

func main() {
	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}
	_ = xys

	err = plotData("out.png", xys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

type xy struct{ x, y float64 }

func readData(path string) ([]xy, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys []xy
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
		}
		xys = append(xys, xy{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
}

func plotData(path string, xys []xy) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}
	wt, err := p.WriterTo(512, 512, "png")
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
