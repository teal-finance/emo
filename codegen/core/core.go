package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Ref struct {
	Name    string `json:"name"`
	Emoji   string `json:"emoji"`
	IsError bool   `json:"isError"`
}

func GetRef() []Ref {
	exPath, _ := filepath.Abs("./codegen/ref.json")
	fmt.Println("Open", exPath)
	file, _ := ioutil.ReadFile(exPath)
	data := []Ref{}
	_ = json.Unmarshal([]byte(file), &data)
	//fmt.Println("Data", data)
	return data
}
