package emo

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"

	color "github.com/acmacalister/skittles"
)

// ObjectInfo : print debug info about something.
func ObjectInfo(args ...any) {
	msg := "[" + color.Yellow("object info") + "] "
	for _, a := range args {
		log.Print(msg+"Type: %T Value: %#v", a, a)
	}
}

// Logger : base emo logger.
type Logger struct {
	Name  string
	Print bool
	Color bool
	Skip  SkipEnum
	Hook  func(Event)
}

type SkipEnum int

const (
	SkipAuto SkipEnum = 0
	SkipNo   SkipEnum = -1
	SkipYes  SkipEnum = 1
)

// NewLogger : create a logger.
func NewLogger(name string, args ...bool) Logger {
	return NewLoggerWithHook(name, nil, args...)
}

// NewLoggerWithHook : create a Logger with a function hook
func NewLoggerWithHook(name string, hook func(Event), args ...bool) Logger {
	print, color, skip := optional(args)
	return Logger{
		Name:  name,
		Print: print,
		Color: color,
		Skip:  skip,
		Hook:  hook,
	}
}

// S forces the print of the caller function and file:line.
// The optional parameter allows to select the position within the stack.
func (l Logger) S(skip ...int) Logger {
	l.Skip = SkipYes
	if len(skip) > 0 {
		l.Skip = SkipEnum(skip[0])
	}
	return l
}

// S forces the log to be printed.
func (l Logger) P() Logger {
	l.Print = true
	return l
}

// Event : base emo event.
type Event struct {
	Emoji   string
	Log     Logger
	IsError bool
	Args    []any
	From    string
	File    string
	Line    int
	cache   string
}

// Error is the implementation of the error interface.
func (e *Event) Error() string {
	if e.cache == "" {
		text := make([]string, 0, len(e.Args))
		for _, a := range e.Args {
			text = append(text, fmt.Sprintf("%v", a))
		}
		e.cache = strings.Join(text, " ")

		if e.From != "" {
			e.cache += " from " + e.Log.bold(e.From) +
				" in " + e.File + ":" +
				e.Log.white(strconv.Itoa(e.Line))
		}
	}
	return e.cache
}

// Stack forces the extraction and print of the call stack info:
// caller function, file and line number.
// The optional parameter allows to change the position within the call stack.
func (e Event) Stack(level ...int) Event {
	skip := int(e.Log.Skip)
	if len(level) > 0 {
		skip += level[0]
	}

	pc := make([]uintptr, 1)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])
	e.File, e.Line = f.FileLine(pc[0])
	e.From = f.Name()
	return e
}

// Err converts an Event to a standard Go error.
func (e Event) Err() error {
	return &e
}

func (e Event) message() string {
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

func optional(args []bool) (bool, bool, SkipEnum) {
	switch len(args) {
	case 0:
		return true, true, SkipAuto
	case 1:
		return args[0], true, SkipAuto
	case 2:
		return args[0], args[1], SkipAuto
	default:
		fl := SkipNo
		if args[2] {
			fl = SkipYes
		}
		return args[0], args[1], fl
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

func processEvent(emoji string, l Logger, isError bool, args []any) Event {
	e := newEvent(emoji, l, isError, args)

	if isError || l.Print {
		log.Print(e.message())
	}

	if l.Hook != nil {
		l.Hook(e)
	}

	return e
}

func newEvent(emoji string, l Logger, isError bool, args []any) Event {
	e := Event{
		Log:     l,
		Emoji:   emoji,
		IsError: isError,
		Args:    args,
		From:    "",
		File:    "",
		Line:    0,
	}

	if (l.Skip == SkipAuto && e.IsError && l.Print) || int(l.Skip) > 0 {
		e = e.Stack(4)
	}

	return e
}
