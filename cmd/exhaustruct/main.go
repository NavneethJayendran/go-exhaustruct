package main

import (
	"flag"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/NavneethJayendran/go-exhaustruct/v3/analyzer"
)

func main() {
	flag.Bool("unsafeptr", false, "")

	a, err := analyzer.NewAnalyzer(nil, nil, false)
	if err != nil {
		panic(err)
	}

	singlechecker.Main(a)
}
