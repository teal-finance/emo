package main

import (
	"fmt"

	"github.com/teal-finance/emo"
)

func hook(err emo.Error) {
	fmt.Println("Event hook", err)
}

func main() {
	log := emo.NewLoggerWithHook("example", hook)
	log.Info("Test info")
}
