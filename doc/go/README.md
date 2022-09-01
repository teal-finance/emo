# Emo Go library

Emoji based semantic scoped debuging for Go

## Usage

Declare a debug zone in a package:

```go
import (
  "github.com/teal-finance/emo"
)

var zone = emo.NewZone("myLib")
```

Create an event for this zone:

```go
zone.Info("An info message")
```

Output:

> `[myLib] ℹ️  An info message`

## Errors

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

## Enable or disable a zone

To only print the errors:

```go
var zone = emo.NewZone("api", false)
```

Setting that second parameter to `false` will disable the printing for logs that are not errors.

## Timestamp

To add a timestamp within the event message, use the third parameter:

```go
const print = false // default true
const date  = true  // default false
var zone = emo.NewZone("api", print, date)
```

## Color

The event message is by default colorized.
But in some cases this is disturbing, especially when testing.
To disable the color, use the fourth parameter:

```go
const print = false // default true
const date  = true  // default false
const color = false // default true
var zone = emo.NewZone("api", print, date, color)
```

## Call stack info

By default emo prints the call stack info only when `isError` and print is enabled.
To always print it, or to never print it, use the fifth parameter:

```go
const print = false // default true
const date  = true  // default false
const color = false // default true
const stack = false // default auto
var zone = emo.NewZone("api", print, date, color, stack)
```

## Hooks

A callback can be passed to a zone.
It will be executed each time an event is fired:

```go
func hook(evt emo.Event) {
    fmt.Println("Event hook", evt.Error)
}

zone := emo.NewZoneWithHook("example", hook)
zone.Debug("Test msg")
```

## `Event`

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

Note: The fields `From`, `File` and `Line`
are always computed when the `zone.Hook` is set,
even is `stack=false` is set.

## Convert an `Event` to a Go standard error

Often the need is to print an error and to return it as a Go standard error.
In that case, the function `Err()` converts an `Event` to a Go standard error:

```go
func foo(n int) error {
    if n < 0 {
        evt := zone.ParamError("Parameter n must be positive, but got:", n)
        return evt.Err()
    }
    return nil
}
```

## Force printing an event

Sometimes, an non-error event should be still printed, even when the `Zone` is configured with `Zone.Print=false`.
In that case, the `P()` helper function can be used:

```go
var prod  = true
var print = !prod
var zone  = emo.NewZone("api", print)

func start() {
    zone.P().Info("Starting...")
    // ...
}
```

The `P(bool)` function accept an optional parameter to explicitly enable/disable the printing of this event:

```go
func foo(n int) error {
    if n < 0 {
        return zone.P(false).ParamError("Parameter n must be positive, but got:", n).Err()
    }
    return nil
}
```

## Force printing the call stack info

Similarly to `P()`, the `S()` helper function controls the call stack info of the current event:

```go
zone.S().Debug("v=", v)     // always print the call stack info
zone.S(-1).Error("v=", v)   // never print the call stack info
zone.S(0).Error("v=", v)    // only when zone.Print=true (default mode)
zone.S(2).Error("v=", v)    // always print, but use the caller layer one level up
```
