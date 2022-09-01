package main

import (
	"fmt"

	"github.com/teal-finance/emo"
)

func hook(evt emo.Event) {
	fmt.Println("Event hook", evt.Error)
}

func main() {
	log := emo.NewLoggerWithHook("example", hook)
	log.Info("Test info")
}
