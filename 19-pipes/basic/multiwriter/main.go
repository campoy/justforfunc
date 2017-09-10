package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buf := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, os.Stderr, buf)

	fmt.Fprintln(mw, "hello")
	fmt.Println("from buffer: ", buf)
}
