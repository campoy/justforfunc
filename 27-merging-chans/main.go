package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func main() {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	for v := range mergeReflect(a, b, c) {
		fmt.Println(v)
	}
}

func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			go func(c <-chan int) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}

		wg.Wait()
		close(out)
	}()
	return out
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func mergeReflect(chans ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface().(int)
		}
	}()
	return out

}
