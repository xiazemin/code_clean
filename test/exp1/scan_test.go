package main

import (
	"fmt"
	"go/token"
	"testing"
	"text/scanner"
)

func TestScanner(t *testing.T) {
	src := []byte(`package main
import "fmt"
//comment
func main() {
  fmt.Println("Hello, world!")
}
`)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("./test.txt", fset.Base(), len(src))
	s.Init(file, src, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}
