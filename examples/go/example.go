package main

import (
	"fmt"

	"github.com/teal-finance/emo"
)

func hook(evt emo.Event) {
	fmt.Printf("hook has been triggered with event: %+v"+"\n", evt)
}

func main() {
	zone := emo.NewZoneWithHook("example", hook)
	zone.Info("Test info")
}
