package emo

import "testing"

func TestBaseEmo(t *testing.T) {
	em := NewZone("testzone")
	evt := em.Info("infomsg")
	if evt.Error.Error() != "infomsg" {
		t.Fatal(evt.Error.Error(), "!=", "infomsg")
	}
}

func TestEmoTypes(t *testing.T) {
	em := NewZone("testzone")
	evt := em.Info("infomsg")
	if evt.Emoji != "ℹ️" {
		t.Fatal(evt.Emoji, "!=", "ℹ️")
	}
}
