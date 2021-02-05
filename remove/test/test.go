package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("not input dir . for default")
	}

	dir := os.Args[1]
	scanFiles(dir)

}

func scanFiles(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
		if f.IsDir() {
			if f.Name() == "." || f.Name() == ".." {
				continue
			}
			scanFiles(dir + "/" + f.Name())
		} else {
			handle(dir + "/" + f.Name())
		}
	}
}

func handle(name string) {
	if len(name) < 4 {
		return
	}
	if len(name) > 8 && name[len(name)-8:] == "_test.go" {
		return
	}
	if name[len(name)-3:] != ".go" {
		return
	}

	src, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fset, f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(buffer.String())
	ioutil.WriteFile(name, buffer.Bytes(), os.FileMode(0777))
}
