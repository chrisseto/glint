package pkg_test

import (
	"testing"

	"github.com/chrisseto/gatch/pkg"
	"golang.org/x/tools/go/analysis/analysistest"
	// "golang.org/x/tools/go/analysis/passes/printf"
)

func Test(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), pkg.Analyzer, "in")
}
