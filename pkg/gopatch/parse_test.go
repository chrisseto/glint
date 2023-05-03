package gopatch_test

import (
	// "go/ast"
	"go/ast"
	"go/parser"
	"go/token"

	// "go/token"
	"os"
	"testing"

	"github.com/chrisseto/gatch/pkg/gopatch"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/ast/astutil"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/pattern"
)


func TestParse(t *testing.T) {
	f, err := os.Open("testdata/s1028.patch")
	require.NoError(t, err)

	defer f.Close()

	patch, err := gopatch.Parse(f)
	require.NoError(t, err)

	for _, p := range patch.Patches {
		expr, err := parser.ParseExpr(p.Diff.Before())
		pn := pattern.ASTToNode(astutil.Apply(expr, func(c *astutil.Cursor) bool {
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

		t.Logf("Before:\n%s\n\n", p.Diff.Before())
		t.Logf("%#v\n", err)
		t.Logf("%#v\n", expr)
		t.Logf("%#v\n", pn)

		expr, err = parser.ParseExpr(p.Diff.After())
		t.Logf("After:\n%s\n\n", p.Diff.After())
		t.Logf("%#v\n", err)
		t.Logf("%#v\n", expr)

	}
}
