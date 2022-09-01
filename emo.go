package emo

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"

	color "github.com/acmacalister/skittles"
)

// Logger : base emo logger.
type Logger struct {
	Name  string
	Print bool
	Hook  func(Event)
}

// Error : base emo event.
type Error struct {
	Emoji   string
	Log     Logger
	IsError bool
	Args    []any
	From    string
	File    string
	Line    int
}

// NewLogger : create a logger.
func NewLogger(name string, print ...bool) Logger {
	p := true
	if len(print) > 0 {
		p = print[0]
	}
	return Logger{
		Name:  name,
		Print: p,
	}
}

// NewLoggerWithHook : create a Logger with a function hook
func NewLoggerWithHook(name string, hook func(Event), print ...bool) Logger {
	p := true
	if len(print) > 0 {
		p = print[0]
	}
	return Logger{
		Name:  name,
		Print: p,
		Hook:  hook,
	}
}

// ObjectInfo : print debug info about something.
func ObjectInfo(args ...any) {
	msg := "[" + color.Yellow("object info") + "] "
	for _, a := range args {
		log.Print(msg+"Type: %T Value: %#v", a, a)
	}
}

func processEvent(emoji string, l Logger, isError bool, args []any) Error {
	e := newError(emoji, l, isError, args)

	if isError || l.Print {
		log.Print(e.message())
	}

	if l.Hook != nil {
		l.Hook(e)
	}

	return e
}

func newError(emoji string, l Logger, isError bool, args []any) Error {
	pc := make([]uintptr, 1)
	runtime.Callers(4, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	return Error{
		Log:     l,
		Emoji:   emoji,
		IsError: isError,
		Args:    args,
		From:    f.Name(),
		File:    file,
		Line:    line,
	}
}

func (e Error) Error() string {
	text := make([]string, 0, len(e.Args))
	for _, a := range e.Args {
		text = append(text, fmt.Sprintf("%v", a))
	}
	str := strings.Join(text, " ")

	if e.IsError && e.Log.Print {
		str += " from " + color.BoldWhite(e.From) +
			" in " + e.File + ":" +
			color.White(strconv.Itoa(e.Line))
	}

	return str
}

func (e Error) message() string {
	msg := ""
	if e.Log.Name != "" {
		msg += "[" + color.Yellow(e.Log.Name) + "] "
	}

	if e.IsError {
		msg += color.Red("Error") + " "
	}

	msg += e.Emoji + "  " + e.Error()

	return msg
}
