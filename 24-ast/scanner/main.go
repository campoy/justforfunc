package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage:\n\t%s [files]\n", os.Args[0])
		os.Exit(1)
	}

	counts := make(map[string]int)

	fs := token.NewFileSet()

	for _, arg := range os.Args[1:] {
		b, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}

		f := fs.AddFile(arg, fs.Base(), len(b))
		var s scanner.Scanner
		s.Init(f, b, nil, scanner.ScanComments)

		for {
			_, tok, lit := s.Scan()
			if tok == token.EOF {
				break
			}
			if tok == token.IDENT {
				counts[lit]++
			}
		}
	}

	type pair struct {
		s string
		n int
	}
	pairs := make([]pair, 0, len(counts))
	for s, n := range counts {
		if len(s) >= 3 {
			pairs = append(pairs, pair{s, n})
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].n > pairs[j].n })

	for i := 0; i < len(pairs) && i < 5; i++ {
		fmt.Printf("%6d %s\n", pairs[i].n, pairs[i].s)
	}
}
