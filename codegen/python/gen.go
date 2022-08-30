package py

import (
	"github.com/teal-finance/emo/codegen/core"
)

func GenPy(ref []core.Ref) {
	code := codeStart
	for _, item := range ref {
		code += genFunc(item.Name, item.Emoji, item.IsError)
	}

	core.Write("lang/python/pyemo/emo_gen.py", code)
}

func genFunc(name, emoji string, isError bool) string {
	name = core.Uncapitalized(name)
	name = core.SnakeCase(name)

	return `
    def ` + name + `(self, *args):
        return self.emo("` + emoji + `", list(args))
`
}
