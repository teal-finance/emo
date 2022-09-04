package emo

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"

	color "github.com/acmacalister/skittles"
)

//go:generate go run ./codegen

// ObjectInfo prints elements list for debugging purpose.
func ObjectInfo(args ...any) {
	msg := "[" + color.Yellow("object info") + "] "
	for _, a := range args {
		fmt.Print(msg+"Type: %T Value: %#v", a, a)
	}
}

// Zone : base emo zone.
type Zone struct {
	Name      string
	Verbose   ParamType // true => always print ; false => print only when isError
	StackInfo ParamType // StackNo => never print the call stack info, StackAuto => only when Verbose=true and isError
	Hook      func(Event)
}

type ParamType int

const (
	Auto ParamType = 0
	No   ParamType = -1
	Yes  ParamType = 1
)

// NewZone : create a zone.
//
// Optional arguments are booleans:
//
// - verbose: true => print all ; false => only isError ; absent (default) => depends on DefaultZone.Verbose
// - stack: true => print the call stack info ; false => only isError ; absent (default) => depends on DefaultZone.Verbose
func NewZone(name string, args ...bool) Zone {
	return NewZoneWithHook(name, nil, args...)
}

// NewZoneWithHook : create a Zone with a function hook
func NewZoneWithHook(name string, hook func(Event), args ...bool) Zone {
	v, si := optional(args)
	return Zone{
		Name:      name,
		Hook:      hook,
		Verbose:   v,
		StackInfo: si,
	}
}

// optional depends on DefaultZone for the default values.
func optional(args []bool) (verbose, stackInfo ParamType) {
	verbose = DefaultZone.Verbose
	stackInfo = DefaultZone.StackInfo

	if len(args) > 0 {
		verbose = No
		if args[0] {
			verbose = Yes
		}
	}

	if len(args) > 1 {
		stackInfo = No
		if args[1] {
			stackInfo = Yes
		}
	}

	return verbose, stackInfo
}

// DefaultZone allows to use emo without having to create a new zone.
// DefaultZone defines the default values.
var DefaultZone = Zone{
	Name:      "",
	Verbose:   Auto,
	StackInfo: Auto,
	Hook:      nil,
}

// GlobalVerbosity controls the verbose mode for all zones having Verbose=Auto.
func GlobalVerbosity(enable bool) {
	DefaultZone.SetVerbosity(enable)
}

// GlobalStackInfo controls the printing of the call stack information for all zones having StackInfo=Auto.
func GlobalStackInfo(enable ...bool) {
	DefaultZone.SetStackInfo(enable...)
}

// SetHook sets the callback function to be triggered by all zones.
func GlobalHook(hook func(Event)) {
	DefaultZone.SetHook(hook)
}

// SetVerbosity controls the verbose mode.
// When used with DefaultZone, it impacts all zones having Verbose=Auto.
func (zone *Zone) SetVerbosity(enable ...bool) {
	zone.Verbose = Auto
	if len(enable) > 0 {
		if enable[0] {
			zone.Verbose = Yes
		} else {
			zone.Verbose = No
		}
	}
}

// SetStackInfo controls the printing of the call stack information.
func (zone *Zone) SetStackInfo(enable ...bool) {
	zone.StackInfo = Auto
	if len(enable) > 0 {
		if enable[0] {
			zone.StackInfo = Yes
		} else {
			zone.StackInfo = No
		}
	}
}

// SetHook sets the callback function to be triggered when an Event occurs.
func (zone *Zone) SetHook(hook func(Event)) {
	zone.Hook = hook
}

// GlobalTimestamp enables the insertion of a date/time timestamp for each event print.
// To disable use `emo.GlobalTimestamp(false)`.
func GlobalTimestamp(enable ...bool) {
	timestampPrefixed = true
	if len(enable) > 0 {
		timestampPrefixed = enable[0]
	}
}

// GlobalColoring enables the coloring of some portion of the event print.
// To disable use `emo.GlobalColoring(false)`.
func GlobalColoring(enable ...bool) {
	outputColored = true
	if len(enable) > 0 {
		outputColored = enable[0]
	}
}

// timestampPrefixed = true => prints are prefixed with the current timestamp
var timestampPrefixed bool = false

// outputColored = true => print colors (use color=false for testing)
var outputColored bool = true

// GlobalTracing enables the print of the Trace() events.
func GlobalTracing(enable bool) {
	tracePrinted = enable
}

var tracePrinted bool = false

// P forces the event to be always printed, even when zone.Verbose=false.
//
// P can also be used to never print the log. Example:
//
//	zone.P(false).Error("my error")
func (zone Zone) P(print ...bool) Zone {
	zone.Verbose = Yes
	if len(print) > 0 && !print[0] {
		zone.Verbose = No
	}
	return zone
}

// S forces the retrieval of the call stack info (caller function and file:line).
// The position within the stack can be changed with the optional parameter.
// S can also be used to disable the call stack info
func (zone Zone) S(skip ...int) Zone {
	zone.StackInfo = Yes
	if len(skip) > 0 {
		zone.StackInfo = ParamType(skip[0])
	}
	return zone
}

// N changes temporary the current event name.
func (zone Zone) N(name string) Zone {
	zone.Name = name
	return zone
}

func (zone Zone) NewEvent(emoji string, isError bool, args ...any) Event {
	e := Event{
		Zone:    zone,
		Emoji:   emoji,
		IsError: isError,
		Args:    args,
		From:    "",
		File:    "",
		Line:    0,
	}

	if e.stackInfoPrinted() {
		e = e.Stack(4)
	}

	return e
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

func (e Event) Print() Event {
	if e.toBePrinted() {
		m := e.Message()
		if timestampPrefixed {
			log.Print(m)
		} else {
			fmt.Println(m)
		}
	}

	return e
}

func (e Event) Message() string {
	msg := ""
	if e.Zone.Name != "" {
		msg += "[" + yellow(e.Zone.Name) + "] "
	}

	msg += e.Emoji

	if e.IsError {
		msg += " " + red("ERROR")
	}

	msg += "  " + e.Error()

	return msg
}

// Err converts an Event to a standard Go error.
func (e Event) Err() error {
	return &e
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
			e.cache += " from " + bold(e.From) +
				" in " + e.File + ":" +
				white(strconv.Itoa(e.Line))
		}
	}
	return e.cache
}

func (zone Zone) hookPresent() bool {
	return (zone.Hook != nil) || (DefaultZone.Hook != nil)
}

func (e Event) CallHook() Event {
	if e.Zone.hookPresent() {
		e = e.Stack()

		if e.Zone.Hook != nil {
			e.Zone.Hook(e)
		}

		if DefaultZone.Hook != nil {
			DefaultZone.Hook(e)
		}
	}
	return e
}

// Stack extracts info from the the call stack info
// (caller function, file and line number).
// Then the event may be printed later.
// The optional parameter allows to change the position within the call stack.
func (e Event) Stack(level ...int) Event {
	// do not compute again the runtime functions if already done
	if e.From == "" {
		skip := int(e.Zone.StackInfo)
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

func (zone Zone) enabled(isError bool) bool {
	return zone.toBePrinted(isError) || zone.hookPresent()
}

func (e Event) toBePrinted() bool {
	return e.Zone.toBePrinted(e.IsError)
}

func (e Event) stackInfoPrinted() bool {
	return e.Zone.stackInfoToBeAppended(e.IsError)
}

func (zone Zone) toBePrinted(isError bool) bool {
	if zone.Verbose >= Yes {
		return true
	}
	if zone.Verbose <= No {
		return false
	}
	if DefaultZone.Verbose >= Yes {
		return true
	}
	if DefaultZone.Verbose <= No {
		return isError
	}
	return true
}

func (zone Zone) stackInfoToBeAppended(isError bool) bool {
	if zone.StackInfo >= Yes {
		return true
	}
	if zone.StackInfo <= No {
		return false
	}
	if DefaultZone.StackInfo >= Yes {
		return true
	}
	if DefaultZone.StackInfo <= No {
		return isError && zone.toBePrinted(isError)
	}
	return false
}

// --- following functions for compatibility with logrus and Go standard logger ---

func (zone Zone) Default() *log.Logger {
	return log.Default()
}

func Trace(args ...any) Event {
	return DefaultZone.Trace(args...)
}

func Tracef(format string, v ...any) Event {
	return DefaultZone.Tracef(format, v...)
}

func Print(args ...any) Event {
	return DefaultZone.Print(args...)
}

func Printf(format string, v ...any) Event {
	return DefaultZone.Printf(format, v...)
}

func Warn(args ...any) Event {
	return DefaultZone.Warn(args...)
}

func Warnf(format string, v ...any) Event {
	return DefaultZone.Warnf(format, v...)
}

func Fatal(args ...any) {
	DefaultZone.Fatal(args...)
}

func Fatalf(format string, v ...any) {
	DefaultZone.Fatalf(format, v...)
}

func Panic(args ...any) {
	DefaultZone.Panic(args...)
}

func Panicf(format string, v ...any) {
	DefaultZone.Panicf(format, v...)
}

func (zone Zone) Trace(args ...any) Event {
	if tracePrinted {
		return zone.NewEvent("👣", false, args...).Print()
	}
	var evt Event
	return evt
}

func (zone Zone) Tracef(format string, v ...any) Event {
	if tracePrinted {
		s := fmt.Sprintf(format, v...)
		return zone.Trace(s)
	}
	var evt Event
	return evt
}

func (zone Zone) Print(args ...any) Event {
	return zone.P().NewEvent("📰", false, args...).Print().CallHook()
}

func (zone Zone) Printf(format string, v ...any) Event {
	s := fmt.Sprintf(format, v...)
	return zone.Print(s)
}

// Warn looks same as Warning, but:
//
// - Warn is printed even when Verbose=No, and
// - Warn is printed with the call-stack info (similar to Error).
func (zone Zone) Warn(args ...any) Event {
	return zone.P().S().Warning(args...)
}

func (zone Zone) Warnf(format string, v ...any) Event {
	return zone.P().S().Warningf(format, v...)
}

func (zone Zone) Fatal(args ...any) {
	m := zone.NewEvent("🤯", true, args...).Message()
	log.Fatal(m)
}

func (zone Zone) Fatalf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	m := zone.NewEvent("🤯", true, s).Message()
	log.Fatal(m)
}

func (zone Zone) Panic(args ...any) {
	m := zone.NewEvent("😵", true, args...).Message()
	log.Panic(m)
}

func (zone Zone) Panicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	m := zone.NewEvent("😵", true, s).Message()
	log.Panic(m)
}

// --- following functions colorize output depending on configuration ---

func red(s string) string {
	if outputColored {
		return color.Red(s)
	}
	return s
}

func yellow(s string) string {
	if outputColored {
		return color.Yellow(s)
	}
	return s
}

func bold(s string) string {
	if outputColored {
		return color.BoldWhite(s)
	}
	return s
}

func white(s string) string {
	if outputColored {
		return color.White(s)
	}
	return s
}
