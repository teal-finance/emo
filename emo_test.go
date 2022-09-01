package emo

import (
	"fmt"
	"strings"
	"testing"
)

func TestBaseEmo(t *testing.T) {
	zone := NewZone("testLogger", true, false)
	evt := zone.Info("info msg")
	if evt.Error() != "info msg" {
		t.Error("got:", evt.Error(), " want:", "info msg")
	}
}

func TestEmoTypes(t *testing.T) {
	zone := NewZone("testLogger", true, false)
	evt := zone.Info("info msg")
	if evt.Emoji != "ℹ️" {
		t.Error("got:", evt.Emoji, " want:", "ℹ️")
	}
}

func TestEmo_Error(t *testing.T) {
	zone := NewZone("testLogger", true, false)
	err := zone.Info("info msg").Stack(0).Err()
	str := fmt.Sprint(err)
	if err.Error() != str {
		t.Error("fmt.Sprint(err) = ", str)
		t.Error("err.Error()     = ", err.Error())
	}
}

func TestEmo_ComputeFileLine(t *testing.T) {
	zone := NewZone("testLogger", true, false)
	err := zone.Info("info msg").Stack().Err()
	prefix := "info msg from "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Error("got----------->", err)
		t.Error("want prefix -->", prefix)
	}
}

func TestEmo_S(t *testing.T) {
	zone := NewZone("testLogger", true, false)
	err := zone.S().Info("info msg").Err()
	prefix := "info msg from "
	if !strings.HasPrefix(err.Error(), prefix) {
		t.Error("got----------->", err)
		t.Error("want prefix -->", prefix)
	}
}
