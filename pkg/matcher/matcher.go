package matcher

import (
	// "fmt"
	// "bytes"
	"fmt"
	"go/ast"
	"go/types"

	// "reflect"

	// "github.com/chrisseto/gatch/pkg/gopatch"
	"golang.org/x/tools/go/ast/inspector"
	// "honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/pattern"
)

type Context = pattern.Matcher

type Pattern interface {
	Matches(Context, ast.Node) bool
	Nodes() []ast.Node
}

type Matcher struct {
	inspector *inspector.Inspector
	types     *types.Info
}

func (m *Matcher) Find() []Match {
	var matches []Match

	// matcher := pattern.Matcher{
	// 	TypesInfo: m.types,
	// 	State:     make(map[string]any),
	// }

	// pattern.Pattern{
	// 	Root: ,
	// }

	// matcher.Match()

	// p := pattern.MustParse(`(CallExpr (Symbol "fmt.Sprintf") (Binding "args" _))`)
	p := pattern.MustParse(`
	(CallExpr (Symbol "time.Now") (Binding "args" _))

	(CallExpr
		
		["subtracted"@(_)]
	)
`)

	fmt.Printf("%#v\n", p)

	// p := pattern.Pattern{
	// 	Root: pattern.CallExpr{
	// 		Fun: pattern.Symbol{
	// 			Name: pattern.String("fmt.Sprintf"),
	// 		},
	// 		Args: pattern.Any{},
	// 	},
	// 	Relevant: make(map[reflect.Type]struct{}),
	// }

	m.inspector.Preorder([]ast.Node{
		(*ast.CallExpr)(nil),
	}, func(n ast.Node) {
		matcher := pattern.Matcher{
			TypesInfo: m.types,
			State:     make(map[string]any),
		}

		if matcher.Match(p, n) {
			fmt.Printf("%#v\n", matcher)
			matches = append(matches, &match{
				nodes: []ast.Node{n},
			})
		}
	})
	return matches
}
