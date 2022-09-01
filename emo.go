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
	Color bool
	Hook  func(Error)
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
	cache   string
}

// NewLogger : create a logger.
func NewLogger(name string, args ...bool) Logger {
	return NewLoggerWithHook(name, nil, args...)
}

// NewLoggerWithHook : create a Logger with a function hook
func NewLoggerWithHook(name string, hook func(Error), args ...bool) Logger {
	print, color := optional(args)
	return Logger{
		Name:  name,
		Print: print,
		Color: color,
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

func (e *Error) Error() string {
	if e.cache == "" {
		text := make([]string, 0, len(e.Args))
		for _, a := range e.Args {
			text = append(text, fmt.Sprintf("%v", a))
		}
		e.cache = strings.Join(text, " ")

		if e.IsError && e.Log.Print {
			e.cache += " from " + e.Log.bold(e.From) +
				" in " + e.File + ":" +
				e.Log.white(strconv.Itoa(e.Line))
		}
	}
	return e.cache
}

func (e Error) message() string {
	msg := ""
	if e.Log.Name != "" {
		msg += "[" + e.Log.yellow(e.Log.Name) + "] "
	}

	if e.IsError {
		msg += e.Log.red("Error") + " "
	}

	msg += e.Emoji + "  " + e.Error()

	return msg
}

func optional(args []bool) (bool, bool) {
	switch len(args) {
	case 0:
		return true, true
	case 1:
		return args[0], true
	default:
		return args[0], args[1]
	}
}

func (l Logger) red(s string) string {
	if l.Color {
		return color.Red(s)
	}
	return s
}

func (l Logger) yellow(s string) string {
	if l.Color {
		return color.Yellow(s)
	}
	return s
}

func (l Logger) bold(s string) string {
	if l.Color {
		return color.BoldWhite(s)
	}
	return s
}

func (l Logger) white(s string) string {
	if l.Color {
		return color.White(s)
	}
	return s
}
