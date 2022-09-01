package emo

import (
	"errors"
	"fmt"
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

// Event : base emo event.
type Event struct {
	Error   error
	Emoji   string
	From    string
	File    string
	Log     Logger
	Line    int
	IsError bool
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
		fmt.Println(msg+"Type: %T Value: %#v", a, a)
	}
}

func processEvent(emoji string, l Logger, isError bool, args []any) Event {
	event := newEvent(emoji, l, isError, args)

	if isError || l.Print {
		fmt.Println(event.message())
	}

	if l.Hook != nil {
		l.Hook(event)
	}

	return event
}

func newEvent(emoji string, l Logger, isError bool, args []any) Event {
	pc := make([]uintptr, 1)
	runtime.Callers(4, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	return Event{
		Log:     l,
		Emoji:   emoji,
		IsError: isError,
		Error:   concatenateErrors(args),
		From:    f.Name(),
		File:    file,
		Line:    line,
	}
}

func concatenateErrors(args []any) error {
	texts := []string{}

	for _, a := range args {
		str := fmt.Sprintf("%v", a)
		texts = append(texts, str)
	}

	all := strings.Join(texts, " ")

	return errors.New(all)
}

func (event Event) message() string {
	msg := ""
	if event.Log.Name != "" {
		msg += "[" + color.Yellow(event.Log.Name) + "] "
	}

	if event.IsError {
		msg += color.Red("Error") + " "
	}

	msg += event.Emoji + "  " + event.Error.Error()

	if event.IsError && event.Log.Print {
		msg += " from " + color.BoldWhite(event.From) +
			" in " + event.File + ":" +
			color.White(strconv.Itoa(event.Line))
	}

	return msg
}
