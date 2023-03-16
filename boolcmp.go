package boolcmp

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const doc = "boolcmp is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name:     "boolcmp",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{},
}

func run(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			if n == nil {
				return false
			}
			switch x := n.(type) {
			case *ast.BinaryExpr:
				if isComparison(x.Op) {
					if lhs := pass.TypesInfo.Types[x.X].Type; lhs != nil {
						if lhs.Underlying().String() == "bool" {
							pass.Reportf(x.Pos(), "bool value used in comparison")
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func isComparison(op token.Token) bool {
	switch op {
	case token.EQL, token.NEQ:
		return true
	default:
		return false
	}
}
