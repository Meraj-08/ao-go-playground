package store

import "testing"

func TestNormalizeLabel(t *testing.T) {
	if got := NormalizeLabel("  Urgent "); got != "urgent" {
		t.Fatalf("got %q", got)
	}
}
