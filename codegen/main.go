package main

import (
	"flag"
	"fmt"

	"github.com/teal-finance/emo/codegen/core"
	"github.com/teal-finance/emo/codegen/dart"
	"github.com/teal-finance/emo/codegen/doc"
	"github.com/teal-finance/emo/codegen/golang"
	py "github.com/teal-finance/emo/codegen/python"
	"github.com/teal-finance/emo/codegen/ts"
)

func main() {
	fmt.Println("[codegen] Generator of emo source code")

	dartf := flag.Bool("dart", false, "generate Dart code")
	docf := flag.Bool("doc", false, "generate the documentation")
	gof := flag.Bool("go", false, "generate Go code")
	pyf := flag.Bool("py", false, "generate Python code")
	tsf := flag.Bool("ts", false, "generate Typescript code")
	flag.Parse()

	hasFlag := *dartf || *docf || *gof || *pyf || *tsf
	enableAll := !hasFlag
	if enableAll {
		fmt.Println("[codegen] No flag => generate code for all languages")
	}

	ref := core.GetRef()

	if enableAll || *dartf {
		fmt.Println("[codegen] Generating Dart code")
		dart.GenCode(ref)
	}

	if enableAll || *docf {
		fmt.Println("[codegen] Generating documentation")
		doc.GenDoc(ref)
	}

	if enableAll || *gof {
		fmt.Println("[codegen] Generating Go code")
		golang.GenGo(ref)
	}

	if enableAll || *pyf {
		fmt.Println("[codegen] Generating Python code")
		py.GenPy(ref)
	}

	if enableAll || *tsf {
		fmt.Println("[codegen] Generating Typescript code")
		ts.GenTs(ref)
	}
}
