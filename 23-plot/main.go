package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	rand.Seed(time.Now().Unix())

	var s server

	http.HandleFunc("/", s.root)
	http.HandleFunc("/statz", errorHandler(s.statz))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type server struct {
	data []time.Duration
	sync.RWMutex
}

func (s *server) root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	x := 1000 * rand.Float64()
	d := time.Duration(x) * time.Millisecond
	// time.Sleep(d)
	fmt.Fprintln(w, "slept for", d)

	s.Lock()
	s.data = append(s.data, d)
	if len(s.data) > 1000 {
		s.data = s.data[len(s.data)-1000:]
	}
	s.Unlock()
}

func (s *server) statz(w http.ResponseWriter, r *http.Request) error {
	s.RLock()
	defer s.RUnlock()

	xys := make(plotter.XYs, len(s.data))
	for i, d := range s.data {
		xys[i].X = float64(i)
		xys[i].Y = float64(d) / float64(time.Millisecond)
	}
	sc, err := plotter.NewScatter(xys)
	if err != nil {
		return errors.Wrap(err, "could not create scatter")
	}
	sc.GlyphStyle.Shape = draw.CrossGlyph{}

	avgs := make(plotter.XYs, len(s.data))
	sum := 0.0
	for i, d := range s.data {
		avgs[i].X = float64(i)
		sum += float64(d)
		avgs[i].Y = sum / (float64(i+1) * float64(time.Millisecond))
	}
	l, err := plotter.NewLine(avgs)
	if err != nil {
		return errors.Wrap(err, "could not create line")
	}
	l.Color = color.RGBA{G: 255, A: 255}

	g := plotter.NewGrid()
	g.Horizontal.Color = color.RGBA{R: 255, A: 255}
	g.Vertical.Width = 0

	p, err := plot.New()
	if err != nil {
		return errors.Wrap(err, "could not create plot")
	}
	p.Add(sc, l, g)
	p.Title.Text = "Endpoint latency"
	p.Y.Label.Text = "ms"
	p.X.Label.Text = "sample"

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return errors.Wrap(err, "could not create writer to")
	}

	w.Header().Set("Content-Type", "image/png")
	_, err = wt.WriteTo(w)
	return errors.Wrap(err, "could not write to output")
}

func errorHandler(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
