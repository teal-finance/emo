package emo

import "testing"

func TestBaseEmo(t *testing.T) {
	log := NewLogger("testLogger")
	err := log.Info("infomsg")
	if err.Error() != "infomsg" {
		t.Fatal(err.Error(), "!=", "infomsg")
	}
}

func TestEmoTypes(t *testing.T) {
	log := NewLogger("testLogger")
	err := log.Info("infomsg")
	if err.Emoji != "ℹ️" {
		t.Fatal(err.Emoji, "!=", "ℹ️")
	}
}
