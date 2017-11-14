package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var s server

	http.HandleFunc("/", s.root)
	http.HandleFunc("/statz", s.statz)
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
	time.Sleep(d)
	fmt.Fprintln(w, "slept for", d)

	s.Lock()
	s.data = append(s.data, d)
	if len(s.data) > 1000 {
		s.data = s.data[len(s.data)-1000:]
	}
	s.Unlock()
}

func (s *server) statz(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()

	for _, d := range s.data {
		fmt.Fprintln(w, d)
	}
}
