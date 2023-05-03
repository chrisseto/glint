package matcher

import (
	"bytes"
	"fmt"
	"go/ast"

	"github.com/chrisseto/glint/pkg/gopatch"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"honnef.co/go/tools/analysis/code"
)

var Analyzer = &analysis.Analyzer{
	Name: "match",
	Doc:  "todo",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run: func(p *analysis.Pass) (any, error) {
		inspect := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)

		patch, err := gopatch.Parse(bytes.NewBuffer([]byte(`
-import "errors"

-errors.New(fmt.Sprintf(...))
+fmt.Errorf(...)
`)))
		if err != nil {
			return nil, err
		}

		before := patch.Patches[0].Diff.Before()
		after := patch.Patches[0].Diff.After()
		fmt.Printf("Before: %s\n", before.Root.String())
		fmt.Printf("After: %s\n", after.Root.String())

		inspect.Preorder([]ast.Node{
			(*ast.CallExpr)(nil),
		}, func(n ast.Node) {
			_, edits, ok := code.MatchAndEdit(p, before, after, n)
			if ok {
				fmt.Printf("Matched! %s\n", edits[0].NewText)
			}
		})

		return nil, nil
	},
}
