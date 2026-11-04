package store

import "testing"

func TestHasTag(t *testing.T) {
	if !HasTag([]string{" Work ", "home"}, "work") {
		t.Fatal("expected a match that ignores case and spacing")
	}
	if HasTag([]string{"home"}, "work") {
		t.Fatal("expected no match")
	}
}
