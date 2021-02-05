package visit

import "go/ast"

type visitor struct {
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.Ident:
		v.use(v.pkg.Info.Uses[node])

	case *ast.ValueSpec:
		if !v.insideFunc {
			for _, ident := range node.Names {
				if !isReserved(ident.Name) {
					v.decl(v.pkg.Info.Defs[ident])
				}
			}
		}
		for _, val := range node.Values {
			ast.Walk(v, val)
		}
		return nil

	case *ast.FuncDecl:
		if node.Body != nil {
			v.insideFunc = true
			ast.Walk(v, node.Body)
			v.insideFunc = false
		}

		return nil
	}

	return v
}
