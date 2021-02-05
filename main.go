package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	src := []byte(`/*comment0*/
package main
import "fmt"
//comment1
/*comment2*/
func main() {
  fmt.Println("Hello, world!")
}
`)

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	//ast.Print(fset, f)

	ast.Inspect(f, func(n ast.Node) bool {
		//find comment
		if comment, ok := n.(*ast.Comment); ok {
			fmt.Println(comment)
			return true
		}

		if commentGroup, ok := n.(*ast.CommentGroup); ok {
			fmt.Println(commentGroup)
			return true
		}
		// Find Return Statements
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			//	t.Logf("return statement found on line %d:\n\t", fset.Position(ret.Pos()).Line)
			printer.Fprint(os.Stdout, fset, ret)
			return true
		}

		// Find Functions
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			var exported string
			if fn.Name.IsExported() {
				exported = "exported "
			}
			fmt.Println(exported)
			//t.Logf("%sfunction declaration found on line %d: %s", exported, fset.Position(fn.Pos()).Line, fn.Name.Name)
			return true
		}

		return true
	})

}
