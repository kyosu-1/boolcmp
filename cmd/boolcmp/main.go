package main

import (
	"github.com/kyosu-1/boolcmp"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(boolcmp.Analyzer) }
