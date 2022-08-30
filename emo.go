package emo

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	color "github.com/acmacalister/skittles"
)

// Zone : base emo zone.
type Zone struct {
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
	Zone    Zone
	Line    int
	IsError bool
}

// NewZone : create a logger.
func NewZone(name string, print ...bool) Zone {
	p := true
	if len(print) > 0 {
		p = print[0]
	}
	return Zone{
		Name:  name,
		Print: p,
	}
}

// NewZoneWithHook : create a zone constructor with a hook
func NewZoneWithHook(name string, hook func(Event), print ...bool) Zone {
	p := true
	if len(print) > 0 {
		p = print[0]
	}
	return Zone{
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

func processEvent(emoji string, zone Zone, isError bool, args []any) Event {
	event := newEvent(emoji, zone, isError, args)

	if isError || zone.Print {
		fmt.Println(event.message())
	}

	if zone.Hook != nil {
		zone.Hook(event)
	}

	return event
}

func newEvent(emoji string, zone Zone, isError bool, args []any) Event {
	pc := make([]uintptr, 1)
	runtime.Callers(4, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	return Event{
		Zone:    zone,
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
	if event.Zone.Name != "" {
		msg += "[" + color.Yellow(event.Zone.Name) + "] "
	}

	if event.IsError {
		msg += color.Red("Error") + " "
	}

	msg += event.Emoji + "  " + event.Error.Error()

	if event.IsError && event.Zone.Print {
		msg += " from " + color.BoldWhite(event.From) +
			" in " + event.File + ":" +
			color.White(strconv.Itoa(event.Line))
	}

	return msg
}
