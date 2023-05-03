package matcher

import (
	"go/ast"
	"go/types"

	"honnef.co/go/tools/pattern"
)

type Wildcard struct {
}

func (*Wildcard) Matches(types.Object) bool {
	return true
}

type ArgumentPattern struct {
}

type NamedPattern struct {
	Package  string
	Function string
}

func (p *NamedPattern) Nodes() []ast.Node {
	return []ast.Node{
		(*ast.CallExpr)(nil),
	}
}

// func (p *NamedPattern) Matches(ctx Context, n ast.Node) bool {
// 	fn := typeutil.Callee(ctx.Types, n.(*ast.CallExpr))

// 	return fn.Name() == p.Function && fn.Pkg().Path() == p.Package
// }

type CallPattern struct {
	pattern.CallExpr
}

func (p *CallPattern) Nodes() []ast.Node {
	return []ast.Node{
		(*ast.CallExpr)(nil),
	}
}

func (p *CallPattern) Matches(ctx Context, n ast.Node) bool {
	// pattern.CallExpr

	// call := n.(*ast.CallExpr)

	// if !p.Func.Matches(ctx, call) {
	// 	return false
	// }

	// return true

	// if len(p.Arguments) != len(call.Args) {
	// 	return false
	// }

	// for i, arg := range p.Arguments {
	// 	if !arg.Matches(ctx, call.Args[i]) {
	// 		return false
	// 	}
	// }

	return true
}
