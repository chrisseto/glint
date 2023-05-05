package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"os"
	"strings"

	"github.com/chrisseto/glint/pkg/gopatch"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"
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
		// fmt.Printf("Before: %s\n", before.Root.String())
		// fmt.Printf("After: %s\n", after.Root.String())

		inspect.Preorder([]ast.Node{
			(*ast.CallExpr)(nil),
		}, func(n ast.Node) {
			_, edits, ok := code.MatchAndEdit(p, before, after, n)
			if !ok {
				return
			}

			// TODO get messages
			// TODO group into files
			p.Report(analysis.Diagnostic{
				SuggestedFixes: []analysis.SuggestedFix{{
					TextEdits: edits,
				}},
			})
		})

		return nil, nil
	},
}

func main() {
	if err := run(); err != nil {
		_ = errors.New(fmt.Sprintf("foo"))
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	conf := packages.Config{
		Mode:  packages.LoadAllSyntax,
		Tests: true,
	}

	pkgs, err := packages.Load(&conf, "./...")
	if err != nil {
		return err
	}

	fmt.Printf("loaded packages: %#v\n", pkgs)

	type result struct {
		pkg      *packages.Package
		analyzer *analysis.Analyzer
	}

	// TODO preallocate
	results := map[result]any{}
	diagnostics := map[result][]analysis.Diagnostic{}

	// See multichecker.Run's internals.
	// TODO topological sort on the analysis chain.
	for _, analyzer := range []*analysis.Analyzer{
		inspect.Analyzer,
		Analyzer,
	} {
		for _, pkg := range pkgs {
			resultsOf := map[*analysis.Analyzer]any{}
			for _, required := range analyzer.Requires {
				resultsOf[required] = results[result{
					pkg:      pkg,
					analyzer: required,
				}]
			}

			pass := analysis.Pass{
				Analyzer:     analyzer,
				Fset:         pkg.Fset,
				Files:        pkg.Syntax,
				OtherFiles:   pkg.OtherFiles,
				IgnoredFiles: pkg.IgnoredFiles,
				Pkg:          pkg.Types,
				TypesInfo:    pkg.TypesInfo,
				TypesSizes:   pkg.TypesSizes,
				TypeErrors:   pkg.TypeErrors,
				ResultOf:     resultsOf,
				Report: func(d analysis.Diagnostic) {
					diagnostics[result{
						pkg:      pkg,
						analyzer: analyzer,
					}] = append(diagnostics[result{
						pkg:      pkg,
						analyzer: analyzer,
					}], d)
				},
			}

			var err error
			results[result{
				pkg:      pkg,
				analyzer: analyzer,
			}], err = pass.Analyzer.Run(&pass)
			if err != nil {
				return err
			}
		}
	}

	var diff strings.Builder
	for result, ds := range diagnostics {
		for _, d := range ds {
			for _, fix := range d.SuggestedFixes {
				for _, edit := range fix.TextEdits {
					pos := result.pkg.Fset.Position(edit.Pos)
					var file *ast.File
					for _, file = range result.pkg.Syntax {
						if edit.Pos > file.Pos() && edit.End <= file.End() {
							break
						}
					}

					fmt.Printf("%#v\n", pos)
					fmt.Fprintf(&diff, "--- a%s\n", pos.Filename)
					fmt.Fprintf(&diff, "+++ b%s\n", pos.Filename)
					fmt.Fprintf(&diff, "@@ -%d,1 +%d,1 @@\n", pos.Line, pos.Line)

					var buf strings.Builder
					if err := format.Node(&buf, result.pkg.Fset, file); err != nil {
						return err
					}

					line := strings.Split(buf.String(), "\n")[pos.Line-1]

					fmt.Fprintf(&diff, "-%s\n", line)
					fmt.Fprintf(&diff, "+%s\n", edit.NewText)
				}
			}
		}
	}

	fmt.Printf("diff:\n%s\n", diff.String())

	return nil
}
