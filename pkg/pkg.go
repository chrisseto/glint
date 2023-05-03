package pkg

import (
	// "fmt"

	"github.com/chrisseto/gatch/pkg/matcher"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "fixer",
	Doc:  `Doc`,
	Requires: []*analysis.Analyzer{
		// printf.Analyzer,
		matcher.Analyzer,
	},
	Run: func(p *analysis.Pass) (any, error) {
		// m := p.ResultOf[matcher.Analyzer].(*matcher.Matcher)

		// pattern := &matcher.CallPattern{
		// 	Func: &matcher.NamedPattern{
		// 		Package:  "fmt",
		// 		Function: "Sprintf",
		// 	},
		// 	// Arguments: ,
		// }

		// fmt.Printf("Found: %+v", m.Find())

		return nil, nil
	},
}
