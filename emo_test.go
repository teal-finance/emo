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

	events := []emo.Event{
		z.S().In("S().In"),
		z.S().Inf("S().Inf"),
		emo.DefaultZone.S().In("DZ.S().In"),
		emo.DefaultZone.S().Inf("DZ.S().Inf"),

		z.S().In("S().In"),
		z.S().Inf("S().Inf"),
		emo.DefaultZone.S().In("DZ.S().In"),
		emo.DefaultZone.S().Inf("DZ.S().Inf"),

		z.S().Print("S().Print"),
		z.S().Printf("S().Printf"),
		emo.DefaultZone.S().Print("DZ.S().Print"),
		emo.DefaultZone.S().Printf("DZ.S().Printf"),

		z.S().Print("S().Print"),
		z.S().Printf("S().Printf"),
		emo.DefaultZone.S().Print("DZ.S().Print"),
		emo.DefaultZone.S().Printf("DZ.S().Printf"),

		z.S().Warning("S().Warning"),
		z.S().Warningf("S().Warningf"),
		emo.DefaultZone.S().Warning("DZ.S().Warning"),
		emo.DefaultZone.S().Warningf("DZ.S().Warningf"),

		z.S().Warn("S().Warn"),
		z.S().Warnf("S().Warnf"),
		emo.DefaultZone.S().Warn("DZ.S().Warn"),
		emo.DefaultZone.S().Warnf("DZ.S().Warnf"),

		z.Warn("Warn"),
		z.Warnf("Warnf"),
		emo.DefaultZone.Warn("DZ.Warn"),
		emo.DefaultZone.Warnf("DZ.Warnf"),

		z.S().Error("S().Error"),
		z.S().Errorf("S().Errorf"),
		emo.DefaultZone.S().Error("DZ.S().Error"),
		emo.DefaultZone.S().Errorf("DZ.S().Errorf"),

		z.S(0).Error("S(0).Error"),
		z.S(0).Errorf("S(0).Errorf"),
		emo.DefaultZone.S(0).Error("DZ.S(0).Error"),
		emo.DefaultZone.S(0).Errorf("DZ.S(0).Errorf"),

		emo.DefaultZone.Error("DZ.Error"),
		emo.DefaultZone.Errorf("DZ.Errorf"),
	}

	for i := range events {
		if events[i].File != evt.File {
			t.Errorf("[call stack info] Want same file, but got %s=%q S.Info=%q", events[i].Args[0], events[i].File, evt.File)
		} else {
			d := events[i].Line - evt.Line
			if d < 0 || d > 100 {
				t.Errorf("[call stack info] Want similar line #, but got got %s=%d S.Info=%d", events[i].Args[0], events[i].Line, evt.Line)
			}
		}
	}
}
