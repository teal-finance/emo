package main

import (
	"flag"
	"fmt"

	"github.com/teal-finance/emo/codegen/doc"
	"github.com/teal-finance/emo/codegen/golang"
	"github.com/teal-finance/emo/codegen/ts"
)

func main() {
	tsf := flag.Bool("ts", false, "generate Typescript code")
	gof := flag.Bool("go", false, "generate Go code")
	docf := flag.Bool("doc", false, "generate the documentation")
	flag.Parse()

	if *docf {
		fmt.Println("Generating documentation")
		doc.GenDoc()
		return
	}

	if *tsf {
		fmt.Println("Generating Typescript code")
		ts.GenTs()
		return
	}
	if *gof {
		fmt.Println("Generating Go code")
		golang.GenGo()
		return
	}

	fmt.Println("Generating code for all languages")
	ts.GenTs()
	golang.GenGo()
}
