package main

import (
	"fmt"
	"teal-finance/emo"
)

func hook(evt emo.Event) {
	fmt.Println("Event hook", evt.Error)
}

func main() {
	em := emo.NewZoneWithHook("example", hook)
	em.Info("Test info")
}
