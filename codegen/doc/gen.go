package doc

import (
	"fmt"

	"github.com/teal-finance/emo/codegen/core"
)

func GenDoc(ref []core.Ref) {
	code := fileStart
	for _, item := range ref {
		code += genFunc(item.Name, item.Emoji, item.IsError)
	}

	core.Write("doc/events/README.md", code)
}

var fileStart = `
# Emo event types

| Name          |  Emoji |  IsError |
|---------------|:------:|:--------:|
`

func genFunc(name, emoji string, isError bool) string {
	errStr := " "
	if isError {
		errStr = "✔️"
	}
	return fmt.Sprintf("| %-13s |   "+emoji+"   |     "+errStr+"    |"+"\n", name)
}
