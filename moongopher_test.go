package moongopher_test

import (
	"testing"

	"github.com/townewgokgok/moongopher/v0"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, moongopher.Analyzer, "a")
}
