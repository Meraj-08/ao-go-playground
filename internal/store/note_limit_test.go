package store

import (
	"strings"
	"testing"
)

func TestLimitNote(t *testing.T) {
	if got := LimitNote(strings.Repeat("a", 300)); len(got) != MaxNote {
		t.Fatalf("got %d", len(got))
	}
}
