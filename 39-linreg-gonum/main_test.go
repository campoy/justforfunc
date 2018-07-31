package main

import (
	"log"
	"testing"
)

func BenchmarkComputeCostFloats(b *testing.B) {
	b.StopTimer()
	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}
	x, y := toVecs(xys)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		computeCost(x, y, 0, 0)
	}
}

func BenchmarkComputeCostLoop(b *testing.B) {
	b.StopTimer()
	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		computeCostLoop(xys, 0, 0)
	}
}
