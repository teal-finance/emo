package main

import (
	"flag"
	"fmt"

	"github.com/teal-finance/emo/codegen/golang"
	"github.com/teal-finance/emo/codegen/ts"
)

func main() {
	tsf := flag.Bool("ts", false, "generate Typescript code")
	gof := flag.Bool("go", false, "generate Go code")
	flag.Parse()

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
