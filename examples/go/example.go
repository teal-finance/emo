package main

import (
	"fmt"

	"github.com/teal-finance/emo"
)

func hook(evt emo.Event) {
	fmt.Printf("hook has been triggered with event: %+v"+"\n", evt)
}

func main() {
	zone := emo.NewZone("example").SetHook(hook)
	zone.Info("Test info")
	// Output:
	// hook has been triggered with event: {Emoji:ℹ️ Zone:{Name:example Verbose:0 StackInfo:0 Hook:0x48e9c0} IsError:false Args:[Test info] From:main.main File:*/emo/examples/go/example.go Line:16 cache:}
}
