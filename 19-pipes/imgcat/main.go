package main

import (
	"fmt"
	"os"

	"github.com/campoy/justforfunc/19-pipes/imgcat/imgcat"
	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing paths of images to cat")
		os.Exit(2)
	}

	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not cat %s: %v\n", path, err)
		}
	}
}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "could not open image")
	}
	defer f.Close()

	return imgcat.Copy(os.Stdout, f)
}
