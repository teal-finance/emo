package emo

import (
	"fmt"
	"log"
	"os"
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
		fmt.Println(msg+"Type: %T Value: %#v", a, a)
	}
}

// Zone : base emo zone.
type Zone struct {
	Name      string
	Verbose   ParamType // true => always print ; false => print only when isError
	StackInfo ParamType // StackNo => never print the call stack info, StackAuto => only when Verbose=true and isError
	Hook      func(Event)
}

var (
	maxNameLen = 1
	fmtName    = "%-1s"
	fmtNameNC  = "%-3s" // NC = non-color
)

// NewZone : create a zone.
func NewZone(name string) Zone {
	if maxNameLen < len(name) {
		maxNameLen = len(name)
		fmtName = "%-" + strconv.Itoa(maxNameLen) + "s"
		fmtNameNC = "%-" + strconv.Itoa(maxNameLen+2) + "s"
	}
	return Zone{
		Name:      name,
		Verbose:   Auto,
		StackInfo: Auto,
		Hook:      nil,
	}
}

// DefaultZone allows to use emo without having to create a new zone.
// DefaultZone defines the default values.
var DefaultZone = Zone{
	Name:      "",
	Verbose:   Auto,
	StackInfo: Auto,
	Hook:      nil,
}

type ParamType int

const (
	Auto ParamType = 0
	No   ParamType = -1
	Yes  ParamType = 1
)

// GlobalTimestamp enables the insertion of a date/time timestamp for each event print.
// To disable use:
//
//	emo.GlobalTimestamp(false)
func GlobalTimestamp(enable bool) {
	timestampPrefixed = enable
}

// GlobalColoring enables/disables the coloring of the event print.
func GlobalColoring(enable bool) {
	outputColored = enable
}

// GlobalVerbosity controls the verbose mode for all zones having Verbose=Auto.
func GlobalVerbosity(enable bool) {
	if enable {
		DefaultZone = DefaultZone.V()
	} else {
		DefaultZone = DefaultZone.V(false)
	}
}

// GlobalStackInfo enables/disables the call stack information for all zones having StackInfo=Auto.
func GlobalStackInfo(enable bool) {
	if enable {
		DefaultZone = DefaultZone.S()
	} else {
		DefaultZone = DefaultZone.S(-1)
	}
}

// SetHook sets the callback function to be triggered by all zones.
func GlobalHook(hook func(Event)) {
	DefaultZone = DefaultZone.SetHook(hook)
}

// SetVerbosity controls the verbose mode.
// When used with DefaultZone, it impacts all zones having Verbose=Auto.
// V forces the event to be always printed, even when zone.Verbose=false.
//
// V can also be used to never print the log. Example:
//
//	zone.V(false).Error("my error")
func (zone Zone) V(auto ...bool) Zone {
	zone.Verbose = Yes
	if len(auto) > 0 {
		if auto[0] {
			zone.Verbose = Auto
		} else {
			zone.Verbose = No
		}
	}
	return zone
}

// S controls the printing of the call stack information.
// Without arguments, S forces the retrieval
// of the call stack info (i.e. caller function and file:line).
//
// The call stack depth can be shifted with the optional parameter:
//
//	func myLoggerFunction(s string) { zone.S(2).Error(s) }
//
// S can also be used to disable the call stack info using a negative value:
//
//	zone.S(-1).Error("msg")
//
// The zero value is to use the global settings:
//
//	zone.S(0).Error("msg")
func (zone Zone) S(callDepth ...int) Zone {
	// if StackInfo already shifts the call stack depth
	if zone.StackInfo > Yes {
		if len(callDepth) > 0 {
			if callDepth[0] < 0 {
				zone.StackInfo = No
			} else if callDepth[0] > 0 {
				// increment the call stack depth
				zone.StackInfo += ParamType(callDepth[0])
			}
			// callDepth[0]==0 is ignored because it may use a bad call stack depth
		}
	} else {
		zone.StackInfo = Yes
		if len(callDepth) > 0 {
			if callDepth[0] >= 1 {
				callDepth[0]++ // shift the call stack depth one step more
			}
			zone.StackInfo = ParamType(callDepth[0])
		}
	}
	return zone
}

// SetHook sets the callback function to be triggered when an Event occurs.
func (zone Zone) SetHook(hook func(Event)) Zone {
	zone.Hook = hook
	return zone
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

// N changes zone name and can be used, for example,
// to temporary change the current event name:
//
//	zone.N("foo").Info("bar")
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

	if e.stackInfoEnabled() {
		e = e.Stack()
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
	msg := e.Zone.highlightName() + e.Emoji
	if e.IsError {
		msg += " " + red("ERROR")
	}
	return msg + "  " + e.Error()
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
// The optional parameter allows to change the call stack depth within the call stack.
func (e Event) Stack(optionalPosition ...int) Event {
	// do not extract the call stack info if already done
	if e.From == "" {
		callDepth := 4
		if e.Zone.StackInfo > Yes {
			callDepth = 3 + int(e.Zone.StackInfo)
		}
		if len(optionalPosition) > 0 {
			callDepth += optionalPosition[0]
		}
		pc := make([]uintptr, 1)
		runtime.Callers(callDepth, pc)
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

func (e Event) stackInfoEnabled() bool {
	return e.Zone.stackInfoEnabled(e.IsError)
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

func (zone Zone) stackInfoEnabled(isError bool) bool {
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
		return false
	}
	return isError && zone.toBePrinted(isError)
}

// --- following functions for compatibility with logrus and Go standard logger ---

func (zone Zone) Default() *log.Logger {
	return log.Default()
}

func Trace(args ...any) Event {
	if tracePrinted {
		return DefaultZone.NewEvent("👣", false, args...).Print()
	}
	var evt Event
	return evt
}

func Tracef(format string, v ...any) Event {
	if tracePrinted {
		s := fmt.Sprintf(format, v...)
		return DefaultZone.Trace(s)
	}
	var evt Event
	return evt
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

// Warn looks same as Warning, but:
//   - Warn is printed even when Verbose=No, and
//   - Warn is printed with the call stack info (similar to Error).
func Print(args ...any) Event {
	return DefaultZone.V().NewEvent("📰", false, args...).Print().CallHook()
}

func Printf(format string, v ...any) Event {
	s := fmt.Sprintf(format, v...)
	return DefaultZone.V().NewEvent("📰", false, s).Print().CallHook()
}

func (zone Zone) Print(args ...any) Event {
	return zone.V().NewEvent("📰", false, args...).Print().CallHook()
}

func (zone Zone) Printf(format string, v ...any) Event {
	s := fmt.Sprintf(format, v...)
	return zone.V().NewEvent("📰", false, s).Print().CallHook()
}

// Warn looks same as Warning, but:
//   - Warn is printed even when Verbose=No, and
//   - Warn is printed with the call stack info (similar to Error).
func Warn(args ...any) Event {
	return DefaultZone.V().S(1).Warning(args...)
}

func Warnf(format string, v ...any) Event {
	return DefaultZone.V().S(1).Warningf(format, v...)
}

// Warn looks same as Warning, but:
//   - Warn is printed even when Verbose=No, and
//   - Warn is printed with the call stack info (similar to Error).
func (zone Zone) Warn(args ...any) Event {
	return zone.V().S(1).Warning(args...)
}

func (zone Zone) Warnf(format string, v ...any) Event {
	return zone.V().S(1).Warningf(format, v...)
}

func Fatal(args ...any) {
	DefaultZone.S(1).Fatal(args...)
}

func Fatalf(format string, v ...any) {
	DefaultZone.S(1).Fatalf(format, v...)
}

func (zone Zone) Fatal(args ...any) {
	msg := zone.S().NewEvent("🤯", true, args...).Message()
	if timestampPrefixed {
		log.Fatal(msg)
	} else {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func (zone Zone) Fatalf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	zone.S(1).Fatal(s)
}

func Panic(args ...any) {
	DefaultZone.S(1).Panic(args...)
}

func Panicf(format string, v ...any) {
	DefaultZone.S(1).Panicf(format, v...)
}

func (zone Zone) Panic(args ...any) {
	msg := zone.S().NewEvent("😵", true, args...).Message()
	if timestampPrefixed {
		log.Panic(msg)
	} else {
		panic(msg)
	}
}

func (zone Zone) Panicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	zone.S(1).Panic(s)
}

// --- following functions colorize output depending on configuration ---

func red(s string) string {
	if outputColored {
		return color.Red(s)
	}
	return s
}

func (zone Zone) highlightName() string {
	if zone.Name == "" {
		return ""
	}
	if outputColored {
		return color.Yellow(fmt.Sprintf(fmtName, zone.Name)) + " "
	}
	return fmt.Sprintf(fmtNameNC, "["+zone.Name+"] ")
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
