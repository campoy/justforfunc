// parser parses the go programs in the given paths and prints
// the top five most common names of local variables and variables
// defined at package level.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage:\n\t%s [files]\n", os.Args[0])
		os.Exit(1)
	}

	fs := token.NewFileSet()
	locals, globals := make(map[string]int), make(map[string]int)

	for _, arg := range os.Args[1:] {
		f, err := parser.ParseFile(fs, arg, nil, parser.AllErrors)
		if err != nil {
			log.Printf("could not parse %s: %v", arg, err)
			continue
		}
		v := newVisitor(f)
		ast.Walk(v, f)
		for k, v := range v.locals {
			locals[k] += v
		}
		for k, v := range v.globals {
			globals[k] += v
		}
	}

	fmt.Println("most common local variable names")
	printTopFive(locals)
	fmt.Println("most common global variable names")
	printTopFive(globals)
}

func printTopFive(counts map[string]int) {
	type pair struct {
		s string
		n int
	}
	pairs := make([]pair, 0, len(counts))
	for s, n := range counts {
		pairs = append(pairs, pair{s, n})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].n > pairs[j].n })

	for i := 0; i < len(pairs) && i < 5; i++ {
		fmt.Printf("%6d %s\n", pairs[i].n, pairs[i].s)
	}
}

type visitor struct {
	pkgDecl map[*ast.GenDecl]bool
	locals  map[string]int
	globals map[string]int
}

func newVisitor(f *ast.File) visitor {
	decls := make(map[*ast.GenDecl]bool)
	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			decls[v] = true
		}
	}

	return visitor{
		decls,
		make(map[string]int),
		make(map[string]int),
	}
}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.AssignStmt:
		if d.Tok != token.DEFINE {
			return v
		}
		for _, name := range d.Lhs {
			countLocalIdent(v, name)
		}
	case *ast.RangeStmt:
		countLocalIdent(v, d.Key)
		countLocalIdent(v, d.Value)
	case *ast.FuncDecl:
		for _, param := range d.Type.Params.List {
			for _, name := range param.Names {
				countLocalIdent(v, name)
			}
		}

		if d.Type.Results == nil {
			return v
		}
		for _, result := range d.Type.Results.List {
			for _, name := range result.Names {
				if name.Name != "" {
					countLocalIdent(v, name)
				}
			}
		}
	case *ast.GenDecl:
		if d.Tok != token.VAR {
			return v
		}
		for _, spec := range d.Specs {
			if value, ok := spec.(*ast.ValueSpec); ok {
				for _, name := range value.Names {
					if name.Name == "_" {
						continue
					}
					if v.pkgDecl[d] {
						v.globals[name.Name]++
					} else {
						v.locals[name.Name]++
					}
				}
			}
		}
	}

	return v
}

func countLocalIdent(v visitor, n ast.Node) {
	ident, ok := n.(*ast.Ident)
	if !ok {
		return
	}
	if ident.Name == "_" {
		return
	}
	if ident.Obj != nil && ident.Obj.Pos() == ident.Pos() {
		v.locals[ident.Name]++
	}
}
