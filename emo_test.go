package emo_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/teal-finance/emo"
)

func TestBaseEmo(t *testing.T) {
	emo.GlobalColoring(false)
	zone := emo.NewZone("testLogger").V().S(-1)
	evt := zone.Info("info msg")
	if evt.Error() != "info msg" {
		t.Error("got:", evt.Error(), " want:", "info msg")
	}
}

func TestEmoTypes(t *testing.T) {
	emo.GlobalColoring(false)
	zone := emo.NewZone("testLogger").V().S(-1)
	evt := zone.Info("info msg")
	if evt.Emoji != "ℹ️" {
		t.Error("got:", evt.Emoji, " want:", "ℹ️")
	}
}

func TestEmo_Error(t *testing.T) {
	emo.GlobalColoring(false)
	zone := emo.NewZone("testLogger").V(false).S(-1)
	err := zone.Error("error msg").Stack(0).Err()
	str := fmt.Sprint(err)
	if err.Error() != str {
		t.Error("fmt.Sprint(err) = ", str)
		t.Error("err.Error()     = ", err.Error())
	}
}

func TestEmo_ComputeFileLine(t *testing.T) {
	emo.GlobalColoring(false)
	zone := emo.NewZone("testLogger").V().S()
	err := zone.Info("info msg").Stack().Err()
	prefix := "info msg from "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Error("got----------->", err)
		t.Error("want prefix -->", prefix)
	}
}

func TestEmo_S(t *testing.T) {
	emo.GlobalColoring(false)
	z := emo.NewZone("testLogger").V().S(-1)
	evt := z.S().Info("info msg")
	err := evt.Err()
	prefix := "info msg from "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Error("got----------->", err)
		t.Error("want prefix -->", prefix)
	}

	s := z.S()
	w, x, y, f, d, e, p, g, q := s.Warning(""), s.Warn(""), z.Warn(""), z.Warnf(""), emo.DefaultZone.Warn(""), z.Errorf(""), s.Print(""), s.Printf(""), emo.DefaultZone.S().Print("")

	if w.File != evt.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v S.Info=%v", w.File, evt.File)
	}
	if w.File != x.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v S.Warn=%v", w.File, x.File)
	}
	if w.File != y.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v Warn=%v", w.File, y.File)
	}
	if w.File != f.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v Warnf=%v", w.File, f.File)
	}
	if w.File != d.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v DefaultZone.Warn=%v", w.File, d.File)
	}
	if w.File != e.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v Errorf=%v", w.File, e.File)
	}
	if w.File != p.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v S.Print=%v", w.File, p.File)
	}
	if w.File != g.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v S.Printf=%v", w.File, g.File)
	}
	if w.File != q.File {
		t.Errorf("want call stack file equality, but got S.Warning=%v DefaultZone.S.Print=%v", w.File, q.File)
	}

	if w.Line == evt.Line {
		t.Errorf("want call stack line difference, but got S.Warning=%v == S.Info=%v", w.Line, evt.Line)
	}
	if w.Line != x.Line {
		t.Errorf("want call stack line equality, but got S.Warning=%v S.Warn=%v", w.Line, x.Line)
	}
	if w.Line != y.Line {
		t.Errorf("want call stack line equality, but got S.Warning=%v Warn=%v", w.Line, y.Line)
	}
	if w.Line != f.Line {
		t.Errorf("want call stack line equality, but got S.Warning=%v Warnf=%v", w.Line, f.Line)
	}
	if w.Line != d.Line {
		t.Errorf("want call stack line equality, but got S.Warning=%v DefaultZone.Warn=%v", w.Line, d.Line)
	}
	if w.Line != e.Line {
		t.Errorf("want call stack line equality, but got S.Warning=%v Errorf=%v", w.Line, e.Line)
	}
	if w.Line != p.Line {
		t.Errorf("want call stack line equality, but got S.Warning=%v S.Print=%v", w.Line, p.Line)
	}
	if w.Line != g.Line {
		t.Errorf("want call stack line equality, but got S.Warning=%v S.Printf=%v", w.Line, g.Line)
	}
	if w.Line != q.Line {
		t.Errorf("want call stack line equality, but got S.Warning=%v DefaultZone.S.Print=%v", w.Line, p.Line)
	}
}
