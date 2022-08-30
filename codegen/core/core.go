package core

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

type Ref struct {
	Name    string `json:"name"`
	Emoji   string `json:"emoji"`
	IsError bool   `json:"isError"`
}

func GetRef() []Ref {
	fn, err := filepath.Abs("./codegen/ref.json")
	if err != nil {
		log.Panic(err)
	}
	log.Print("Open referential: ", fn)

	b, err := os.ReadFile(fn)
	if err != nil {
		log.Panic(err)
	}

	ref := []Ref{}
	err = json.Unmarshal(b, &ref)
	if err != nil {
		log.Panic(err)
	}

	return ref
}

func Write(fn, code string) {
	fn, err := filepath.Abs(fn)
	if err != nil {
		log.Panic(err)
	}

	file, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		log.Panic(err)
	}

	n, err := file.Write([]byte(code))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("File: %s (%d bytes)", fn, n)
}

func Uncapitalized(s string) string {
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func SnakeCase(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
