package matcher

import "go/ast"

type Match interface {
	// Nodes returns root [ast.Node]s that were matched.
	Nodes() []ast.Node
}

type match struct {
	nodes []ast.Node
}

func (m *match) Nodes() []ast.Node {
	return m.nodes
}
