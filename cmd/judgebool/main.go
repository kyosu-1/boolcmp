package main

import (
	"github.com/kyosu-1/judgebool"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(judgebool.Analyzer) }
