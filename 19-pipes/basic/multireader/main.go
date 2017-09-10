package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	header := strings.NewReader("<msg>")
	body := strings.NewReader("hello")
	writer := strings.NewReader("</msg>")

	r := io.MultiReader(header, body, writer)
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
}
