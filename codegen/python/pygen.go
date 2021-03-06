package py

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/teal-finance/emo/codegen/core"
)

func toSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func genFunc(name string, emoji string, isError bool) string {
	a := []rune(name)
	a[0] = unicode.ToLower(a[0])
	n := toSnakeCase(string(a))
	s := "    def " + n + "(self, *args):\n"
	s += `        return self.emo("` + emoji + `", list(args))` + "\n"
	return s
}

func GenPy() {
	data := core.GetRef()
	fs := ""
	for _, item := range data {
		f := genFunc(item.Name, item.Emoji, item.IsError)
		fs += "\n" + f
	}
	_template := "# Code generated by codegen/python/pyemo/pygen.go; DO NOT EDIT\n" + TemplateStart + fs

	_filepath, err := filepath.Abs("./lang/python/pyemo/emo_gen.py")
	if err != nil {
		log.Fatalf("filepath error: %s", err)
	}

	file, err := os.OpenFile(_filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	_data := []byte(_template)

	_, err = file.WriteAt(_data, 0)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	log.Printf("File %s created.\n", _filepath)
}
