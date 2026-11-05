package format

import (
	"bytes"
	"strings"
	"testing"
)

func TestMatchesNothing(t *testing.T) {
	var out bytes.Buffer
	if err := Matches(&out, "plumber", nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), `no tasks match "plumber"`) {
		t.Fatalf("got %q", out.String())
	}
}
