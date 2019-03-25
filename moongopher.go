package moongopher

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "moongopher",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "moongopher is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	g := newGenerator(pass)
	g.process()
	return nil, nil
}
