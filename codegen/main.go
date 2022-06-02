package main

import (
	"flag"
	"fmt"

	"github.com/teal-finance/emo/codegen/dart"
	"github.com/teal-finance/emo/codegen/doc"
	"github.com/teal-finance/emo/codegen/golang"
	py "github.com/teal-finance/emo/codegen/python"
	"github.com/teal-finance/emo/codegen/ts"
)

func main() {
	tsf := flag.Bool("ts", false, "generate Typescript code")
	gof := flag.Bool("go", false, "generate Go code")
	pyf := flag.Bool("py", false, "generate Python code")
	dartf := flag.Bool("dart", false, "generate Dart code")
	docf := flag.Bool("doc", false, "generate the documentation")
	flag.Parse()

	hasFlag := false
	if *docf {
		hasFlag = true
		fmt.Println("Generating documentation")
		doc.GenDoc()
		return
	}

	if *tsf {
		hasFlag = true
		fmt.Println("Generating Typescript code")
		ts.GenTs()
		return
	}

	if *pyf {
		hasFlag = true
		fmt.Println("Generating Python code")
		py.GenPy()
		return
	}

	if *dartf {
		hasFlag = true
		fmt.Println("Generating Dart code")
		dart.GenCode()
		return
	}

	if *gof {
		hasFlag = true
		fmt.Println("Generating Go code")
		golang.GenGo()
		return
	}

	if !hasFlag {
		fmt.Println("Generating code for all languages")
		ts.GenTs()
		golang.GenGo()
		py.GenPy()
		dart.GenCode()
	}
}
