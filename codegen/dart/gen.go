package dart

import (
	"github.com/teal-finance/emo/codegen/core"
)

func GenCode(ref []core.Ref) {
	code := codeStart
	for _, item := range ref {
		code += genFunc(item.Name, item.Emoji, item.IsError)
	}
	code += codeEnd

	core.Write("lang/dart/lib/src/debug.dart", code)
}

func genFunc(name, emoji string, isError bool) string {
	name = core.Uncapitalized(name)

	return `
  /// A debug message for ` + name + `
  ///
  /// emoji: ` + emoji + `
  String ` + name + `([dynamic obj, String? domain]) => emo("` + emoji + `", obj, domain);
`
}
