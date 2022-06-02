package main

import (
	"fmt"

	"github.com/teal-finance/emo"
)

func hook(evt emo.Event) {
	fmt.Println("Event hook", evt.Error)
}

func main() {
	em := emo.NewZoneWithHook("example", hook)
	em.Info("Test info")
}
