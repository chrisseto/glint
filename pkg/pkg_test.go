package pkg_test

import (
	"testing"

	"github.com/chrisseto/glint/pkg"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), pkg.Analyzer, "in")
}
