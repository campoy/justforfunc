package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		_, err := fmt.Fprintln(pw, "hello")
		if err != nil {
			panic(err)
		}
	}()

	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		panic(err)
	}
}
