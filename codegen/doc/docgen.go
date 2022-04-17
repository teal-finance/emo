package doc

import (
	"log"
	"os"
	"path/filepath"

	"github.com/teal-finance/emo/codegen/core"
)

func genFunc(name string, emoji string, isError bool) string {
	errStr := ""
	if isError {
		errStr = "✔️"
	}
	return "|   " + name + "     |   " + emoji + "   |     " + errStr + "    |"
}

func GenDoc() {
	data := core.GetRef()
	fs := ""
	for _, item := range data {
		f := genFunc(item.Name, item.Emoji, item.IsError)
		fs += f + "\n"
	}
	_template := TemplateStart + fs

	_filepath, err := filepath.Abs("./doc/events/README.md")
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
