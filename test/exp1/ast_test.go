package  main
import (
	"fmt"
	"go/token"
	"testing"
	"text/scanner"
)
func TestParserAST(t *testing.T) {
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
    f, err := parser.ParseFile(fset, "", src, 0)
    if err != nil {
        panic(err)
    }

    // Print the AST.
	ast.Print(fset, f)
	
	ast.Inspect(f, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			t.Logf("return statement found on line %d:\n\t", fset.Position(ret.Pos()).Line)
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
			t.Logf("%sfunction declaration found on line %d: %s", exported, fset.Position(fn.Pos()).Line, fn.Name.Name)
			return true
		}
	 
		return true
	
}