# Emo

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

See the complete [events list](#event-types)

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

## Event types

| Name       |  Emoji |  IsError |
|------------|:------:|:--------:|
|   Info     |   ℹ️   |          |
|   Warning     |   🔔   |          |
|   Error     |   💢   |     ✔️    |
|   Query     |   🗄️   |          |
|   QueryError     |   🗄️   |     ✔️    |
|   Encrypt     |   🎼"   |          |
|   EncryptError     |   🎼"   |     ✔️    |
|   Decrypt     |   🗝️   |          |
|   DecryptError     |   🗝️   |     ✔️    |
|   Time     |   ⏱️   |          |
|   TimeError     |   ⏱️   |     ✔️    |
|   Param     |   📥   |          |
|   ParamError     |   📥   |     ✔️    |
|   Debug     |   💊   |          |

## Development

### Run the tests

```bash
go test -tags=emo
```

### Generate the code

Run the codegen tools to build up the functions from the `codegen/ref.json` file. 
Build the Go library:

```bash
go run codegen/main.go -go
```

This will regenerate the `emo_gen.go` file

Build the Typescript library:

```bash
go run codegen/main.go -ts
```

This will regenerate the `lang/typescript/src/emo_gen.ts` file

To build all languages at once use no flag