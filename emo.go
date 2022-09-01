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

// Zone : base emo zone.
type Zone struct {
	Name  string
	Print bool      // true => always print ; false => print only when isError
	Date  bool      // true => prefix the log with the current timestamp
	Color bool      // true => print colors (use color=false for testing)
	Skip  StackEnum // StackNo => never print the call stack info, StackAuto => only when Print=true and isError
	Hook  func(Event)
}

type StackEnum int

const (
	StackAuto StackEnum = 0
	StackNo   StackEnum = -1
	StackYes  StackEnum = 1
)

// NewZone : create a zone.
//
// Optional arguments are booleans:
//
// - print=false => print only when isError, default is true
// - date=true => prefix the log with the current timestamp, default is false
// - color=true => print colors (use color=false for testing), default is true
// - stack=absent (default) => print the call stack info when Print=true and isError, true => always, false => never
func NewZone(name string, args ...bool) Zone {
	return NewLoggerWithHook(name, nil, args...)
}

// NewLoggerWithHook : create a Zone with a function hook
func NewLoggerWithHook(name string, hook func(Event), args ...bool) Zone {
	print, date, color, stack := optional(args)
	return Zone{
		Name:  name,
		Print: print,
		Date:  date,
		Color: color,
		Skip:  stack,
		Hook:  hook,
	}
}

// P forces the log to be printed, even when zone.Pint=false.
func (zone Zone) P(print ...bool) Zone {
	zone.Print = true
	if len(print) > 0 {
		zone.Print = print[0]
	}
	return zone
}

// S forces the print of the call stack info (caller function and file:line).
// The optional parameter allows to select the position within the stack.
func (zone Zone) S(skip ...int) Zone {
	zone.Skip = StackYes
	if len(skip) > 0 {
		zone.Skip = StackEnum(skip[0])
	}
	return zone
}

// Event : base emo event.
type Event struct {
	Emoji   string
	Zone    Zone
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
			e.cache += " from " + e.Zone.bold(e.From) +
				" in " + e.File + ":" +
				e.Zone.white(strconv.Itoa(e.Line))
		}
	}
	return e.cache
}

// Stack extracts info from the the call stack info
// (caller function, file and line number)
// in order to print it later.
// The optional parameter allows to change the position within the call stack.
func (e Event) Stack(level ...int) Event {
	// do not compute again the runtime functions if already done
	if e.From == "" {
		skip := int(e.Zone.Skip)
		if len(level) > 0 {
			skip += level[0]
		}

		pc := make([]uintptr, 1)
		runtime.Callers(skip, pc)
		f := runtime.FuncForPC(pc[0])
		e.File, e.Line = f.FileLine(pc[0])
		e.From = f.Name()
	}
	return e
}

// Err converts an Event to a standard Go error.
func (e Event) Err() error {
	return &e
}

func optional(args []bool) (print, date, color bool, stack StackEnum) {
	// default values
	print, date, color, stack = true, false, true, StackAuto

	switch len(args) {
	case 0:
		return print, date, color, stack
	case 1:
		return args[0], date, color, stack
	case 2:
		return args[0], args[1], color, stack
	case 3:
		return args[0], args[1], args[2], stack
	default:
		stack = StackNo
		if args[3] {
			stack = StackYes
		}
		return args[0], args[1], args[2], stack
	}
}

func new(emoji string, zone Zone, isError bool, args []any) Event {
	e := Event{
		Zone:    zone,
		Emoji:   emoji,
		IsError: isError,
		Args:    args,
		From:    "",
		File:    "",
		Line:    0,
	}

	if (zone.Skip >= StackYes) || (zone.Skip == StackAuto && e.IsError && zone.Print) {
		e = e.Stack(4)
	}

	return e
}

func (e Event) print() Event {
	if e.IsError || e.Zone.Print {
		m := e.message()
		if e.Zone.Date {
			log.Print(m)
		} else {
			fmt.Println(m)
		}
	}
	return e
}

func (e Event) callHook() Event {
	if e.Zone.Hook != nil {
		e = e.Stack()
		e.Zone.Hook(e)
	}
	return e
}

func (e Event) message() string {
	msg := ""
	if e.Zone.Name != "" {
		msg += "[" + e.Zone.yellow(e.Zone.Name) + "] "
	}

	msg += e.Emoji

	if e.IsError {
		msg += " " + e.Zone.red("ERROR")
	}

	msg += "  " + e.Error()

	return msg
}

func (zone Zone) red(s string) string {
	if zone.Color {
		return color.Red(s)
	}
	return s
}

func (zone Zone) yellow(s string) string {
	if zone.Color {
		return color.Yellow(s)
	}
	return s
}

func (zone Zone) bold(s string) string {
	if zone.Color {
		return color.BoldWhite(s)
	}
	return s
}

func (zone Zone) white(s string) string {
	if zone.Color {
		return color.White(s)
	}
	return s
}
