package gopatch

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
	"honnef.co/go/tools/pattern"
)

type Elision struct {
	ast.Expr
}

func (Elision) Convert() pattern.Node {
	return pattern.Binding{
		Name: "e",
		Node: pattern.Any{},
	}
}

func (Elision) Pos() token.Pos {
	return 0
}

func (Elision) End() token.Pos {
	return 0
}

func convert(expr ast.Expr) pattern.Node {
	return pattern.ASTToNode(astutil.Apply(expr, func(c *astutil.Cursor) bool {
		_, ok := c.Node().(*ast.BadExpr)
		if !ok {
			return true
		}

		// if p.Diff.Before()[be.From:be.To] != "..." {
		// 	return true
		// }

		c.Replace(Elision{})
		return false
	}, nil))
}
