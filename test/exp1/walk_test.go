package main 
func TestInspectAST(t *testing.T) {
    // Create the AST by parsing src.
    fset := token.NewFileSet() // positions are relative to fset
    f, err := parser.ParseFile(fset, "./main.go", nil, parser.ParseComments)
    if err != nil {
        panic(err)
    }

    ast.Inspect(f, func(n ast.Node) bool {
        // Find Return Statements
        ret, ok := n.(*ast.ReturnStmt)
        if ok {
            fmt.Printf("return statement found on line %v:\n", fset.Position(ret.Pos()))
            printer.Fprint(os.Stdout, fset, ret)
            fmt.Printf("\n")
            return true
        }
        return true
    })
}