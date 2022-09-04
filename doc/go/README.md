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

`[myLib] ℹ️  An info message`

## Errors

Create an error event:

```go
import errors

err := errors.New("PARAM ERROR")
zone.Error("An error has occurred:", err)
```

Output:

`[myLib] 📥  ERROR An error has occurred: PARAM ERROR from main.main in emo/examples/example.go:17`

It prints additional information about the file and the line
if the event is of type error

See the complete [events list](../events/README.md)

## Enable or disable a zone

To only print the errors:

```go
var zone = emo.NewZone("api", false)
```

Setting that second parameter to `false` disables the printing for logs that are not errors.

This can also be later set using:

```go
zone.SetVerbosity(false) // only print when isError
zone.SetVerbosity(true)  // always print any event (except Trace events)
zone.SetVerbosity()      // revert to default: inherit the global settings
```

To change all zones at once:

```go
emo.GlobalVerbosity(false) // verbose mode for all zones inheriting the global settings
emo.GlobalVerbosity(true)  // only isError for all zones inheriting the global settings
```

## Call stack info

By default emo prints the call-stack info only when `isError` and in verbose mode.
The call-stack info contains the caller function, the source file, and line number.

To always print it, or to never print it, use the third parameter:

```go
const verbose = false // default true
const stack = true    // default auto = only when isError
var zone = emo.NewZone("api", verbose, stack)
```

This can also be later set using a similar function as for `zone.SetVerbosity()`:

```go
zone.SetStackInfo(true)  // print the call-stack info for all events
zone.SetStackInfo(false) // never print the call-stack info for any event
zone.SetStackInfo()      // inherits from global settings
```

To control the call-stack info for all zones:

```go
emo.GlobalStackInfo(true)  // print the call-stack info for the zones inheriting global settings
emo.GlobalStackInfo(false) // never print the call-stack for the zones inheriting global settings
emo.GlobalStackInfo()      // default: only when `isError` and in verbose mode
```

## Force printing an event

Sometimes, an non-error event should be still printed, even when the `Zone` is configured with `Zone.Verbose=false`.
In that case, the `P()` helper function can be used:

```go
var prod    = true
var verbose = !prod
var zone    = emo.NewZone("api", verbose)

func start() {
    zone.P().Info("Starting...")
    // ...
}
```

The `P(bool)` function accept an optional parameter to explicitly enable/disable the printing of an event:

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
zone.S().Debug("v=", v)     // always print the call stack info, like S(1)
zone.S(-1).Error("v=", v)   // never print the call stack info
zone.S(0).Error("v=", v)    // inherits the global settings, usually print when zone.Print=true (default)
zone.S(1).Error("v=", v)    // always print the call stack info, like S()
zone.S(2).Error("v=", v)    // always print, but use the caller layer at one level up
```

## Temporary change the `Event` name

The `N()` helper function sets a temporary name for the current event only:

```go
var zone = emo.NewZone("API")
zone.N("GET").Debug("Received a GET request with v=", v)   // replaces [API] -> [GET]
```

Another example:

```go
var zA = emo.NewZone("A")
var zB = emo.NewZone("B")

zA.Info("this Info event is related to the zone A")
zB.Info("this Info event is related to the zone B")

// previous lines can be replaced by:

emo.DefaultZone.N("A").Info("this Info event is related to the zone A")
emo.DefaultZone.N("B").Info("this Info event is related to the zone B")
```

## Timestamp

To timestamp all event messages:

```go
emo.GlobalTimestamp(true)
```

By default, `emo` does not timestamp the events.

## Color

The event messages are by default colorized.
But in some cases this is disturbing, especially when testing.
To disable the color:

```go
emo.GlobalColoring(false)
```

## Convert an `Event` to a Go standard error

Sometimes a function needs to print an error and to return this same string as a Go standard error.
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

## Global zone

All the `emo` functions can be used without having to create a custom `Zone`.
The `emo.DefaultZone` is a pre-configured `Zone` ready to use:

```go
var zone = emo.DefaultZone
var ok = true
zone.Info("my event", ok)
```

## _Printf_ format

The Go implementation of `emo` enables the _Printf_ format. Each function has a corresponding function ending with `f`:

```go
zone.Info("my event", ok)
zone.Infof("my event %v", ok)
```

## Compatibility with the standard Go logger

To use `emo` as a logger the Go implementation of `emo` also implements the standard Go logger functions:

```go
zone.Print("my event", ok)       // not an error but always printed
zone.Printf("my event %v", ok)
zone.Fatal("my event", ok)       // not an error but always printed
zone.Fatalf("my event %v", ok)
zone.Panic("my event", ok)       // not an error but always printed
zone.Panicf("my event %v", ok)
zone.Default()
```

Moreover, the Go implementation of `emo` also implements some of other common functions of loggers like `logrus`:

```go
zone.Trace("my event", ok)       // not an error but always printed
zone.Tracef("my event %v", ok)
zone.Warn("my event", ok)       // not an error but always printed
zone.Warnf("my event %v", ok)
```

## Enable tracing

The `Trace` events can be very verbose.
Thus a specific function controls its enabling:

```go
emo.GlobalTracing(true)
```
