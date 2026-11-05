package store

import "testing"

func TestCleanNote(t *testing.T) {
	if got := CleanNote("  one \n\n two "); got != "one\n two" {
		t.Fatalf("got %q", got)
	}
}
