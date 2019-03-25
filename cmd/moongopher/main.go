package main

import (
	"github.com/townewgokgok/moongopher/v0"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(moongopher.Analyzer)
}
