package golang

import (
	"log"
	"strconv"

	"github.com/dolmen-go/codegen"
	"github.com/teal-finance/emo/codegen/core"
)

func genFunc(name string, emoji string, isError bool) string {
	return `func (zone Zone) ` + name + `(args ...interface{}) Event {
	return processEvent("` + emoji + `", zone, ` + strconv.FormatBool(isError) + `, args)
}
`
}

func GenGo() {
	data := core.GetRef()
	fs := ""
	for _, item := range data {
		f := genFunc(item.Name, item.Emoji, item.IsError)
		fs += f + "\n"
	}
	template := `// Code generated by codegen/golang/codegen.go; DO NOT EDIT.
	package emo
	` + fs

	tmpl := codegen.MustParse(template)
	f := "emo_gen.go"
	if err := tmpl.CreateFile(f, "emo"); err != nil {
		log.Fatal(err)
	}
	log.Printf("File %s created.\n", f)
}
