package doc

import (
	"github.com/teal-finance/emo/codegen/core"
)

func GenDoc(ref []core.Ref) {
	code := TemplateStart
	for _, item := range ref {
		code += genFunc(item.Name, item.Emoji, item.IsError)
	}

	core.Write("doc/events/README.md", code)
}

var TemplateStart = `
# Emo event types

| Name       |  Emoji |  IsError |
|------------|:------:|:--------:|
`

func genFunc(name, emoji string, isError bool) string {
	errStr := ""
	if isError {
		errStr = "✔️"
	}
	return "|   " + name + "     |   " + emoji + "   |     " + errStr + "    |" + "\n"
}
