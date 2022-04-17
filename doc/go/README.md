# Emo Go library

Emoji based semantic scoped debuging for Go

## Usage

Declare a debug zone in a package:

```go
import (
  emolib "github.com/teal-finance/emo"
)

var emo = emolib.NewZone("testzone")
```

Create an event for this zone:

```go
emo.Info("An info message")
```

Output:

> [testzone] ℹ️  infomsg

### Errors

Create an event of error type:

```go
import errors

err := errors.New("PARAM ERROR")
emo.Error("This is an error message", err)
```

Output:

> [testzone] Error 📥  This is a parameter error message PARAM ERROR from main.main in emo/examples/example.go:17

It prints additionnal information about the file and the line if the event is
of type error

See the complete [events list](../events/README.md)

### Enable or disable a zone

To disable events printing for a zone:

```go
var emo = emolib.NewZone("api", false)
```

Setting the second parameter to `false` will disable the printing for this zone

### Hooks

A callback can be passed to a zone. It will be executed each time an event
is fired:

```go
func hook(evt emo.Event) {
  fmt.Println("Event hook", evt.Error)
}

em := emo.NewZoneWithHook("example", hook)
em.Debug("Test msg")
```

Structure of an `Event`:

```go
type Event struct {
  Error   error
  Emoji   string
  From    string
  File    string
  Zone    Zone
  Line    int
  IsError bool
}
```
