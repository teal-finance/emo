# Emo Go library

Emoji based semantic scoped debuging for Go

## Usage

Declare a debug zone in a package:

```go
import (
  "github.com/teal-finance/emo"
)

var zone = emo.NewLogger("myLib")
```

Create an event for this zone:

```go
zone.Info("An info message")
```

Output:

> `[myLib] ℹ️  An info message`

### Errors

Create an event of error type:

```go
import errors

err := errors.New("PARAM ERROR")
zone.Error("An error has occurred:", err)
```

Output:

> `[myLib] 📥 ERROR  An error has occurred: PARAM ERROR from main.main in emo/examples/example.go:17`

It prints additional information about the file and the line
if the event is of type error

See the complete [events list](../events/README.md)

### Enable or disable a zone

To only print the errors:

```go
var zone = emo.NewLogger("api", false)
```

Setting the second parameter to `false` will disable the printing for logs that are not errors.

### Hooks

A callback can be passed to a zone.
It will be executed each time an event is fired:

```go
func hook(evt emo.Event) {
    fmt.Println("Event hook", evt.Error)
}

zone := emo.NewLoggerWithHook("example", hook)
zone.Debug("Test msg")
```

Exported fields of an `Event`:

```go
type Event struct {
    Emoji   string
    Zone    Zone
    IsError bool
    Args    []any
    From    string
    File    string
    Line    int
}
```
