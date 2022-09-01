package golang

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/dolmen-go/codegen"
	"github.com/teal-finance/emo/codegen/core"
)

func GenGo(ref []core.Ref) {
	template := "" +
		`// Code generated by https://github.com/teal-finance/emo/blob/main/codegen/golang/gen.go ; DO NOT EDIT.

	package emo
	`

	for _, item := range ref {
		template += genFunc(item.Name, item.Emoji, item.IsError)
	}

	tmpl := codegen.MustParse(template)

	fn, err := filepath.Abs("generated.go")
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.CreateFile(fn, "emo"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("[codegen] File: " + fn)
}

func genFunc(name, emoji string, isError bool) string {
	if isError {
		return `
	func (zone Zone) ` + name + `(args ...any) Event {
		return processEvent("` + emoji + `", zone, ` + strconv.FormatBool(isError) + `, args)
	}
	`
	}

	return `
func (zone Zone) ` + name + `(args ...any) Event {
	if zone.Print || zone.Hook != nil {
		return processEvent("` + emoji + `", zone, ` + strconv.FormatBool(isError) + `, args)
	}
	return Event{}
}
`
}
