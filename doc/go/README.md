# Emo Go library

Emoji based logger for Go

## Usage

Declare the emo logger in a package:

```go
import (
  "github.com/teal-finance/emo"
)

var log = emo.NewLogger("myLib")
```

Log an info. message:

```go
log.Info("An info message")
```

Output:

> [myLib] ℹ️  An info message

### Errors

Create an event of error type:

```go
import errors

err := errors.New("PARAM ERROR")
log.Error("An error has occurred:", err)
```

Output:

> [myLib] Error 📥  An error has occurred: PARAM ERROR from main.main in emo/examples/example.go:17

It prints additional information about the file and the line
if the event is of type error

See the complete [events list](../events/README.md)

### Enable or disable a logger

To log only the errors:

```go
var log = emo.NewLogger("api", false)
```

Setting the second parameter to `false` will disable the printing for logs that are not errors.

### Hooks

A callback can be passed to a logger.
It will be executed each time an event is fired:

```go
func hook(evt emo.Event) {
  fmt.Println("Event hook", evt.Error)
}

log := emo.NewLoggerWithHook("example", hook)
log.Debug("Test msg")
```

Structure of an `Event`:

```go
type Event struct {
  Error   error
  Emoji   string
  From    string
  File    string
  Log     Logger
  Line    int
  IsError bool
}
```
