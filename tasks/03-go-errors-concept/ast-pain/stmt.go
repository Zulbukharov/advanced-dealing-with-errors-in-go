package astpain

import (
	"go/ast"
)

func Name(expr ast.Expr, name string) string {
	switch tt := expr.(type) {
	case *ast.Ident:
		return name + tt.Name
	case *ast.SelectorExpr:
		switch xx := tt.X.(type) {
		case *ast.Ident:
			return xx.Name + "." + tt.Sel.Name + name
		case *ast.SelectorExpr:
			return Name(xx, name+"."+tt.Sel.Name)
		}
	default:
		return "anonymous func"
	}
	return name
}

// GetDeferredFunctionName возвращает имя функции, вызов которой был отложен через defer,
// если входящий node является *ast.DeferStmt.
func GetDeferredFunctionName(node ast.Node) string {
	if def, ok := node.(*ast.DeferStmt); ok {
		return Name(def.Call.Fun, "")
	}
	return ""
}
