package format

import (
	"strings"
	"testing"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

func TestTable(t *testing.T) {
	var b strings.Builder
	err := Table(&b, []store.Task{{ID: 1, Title: "a", Done: true}, {ID: 2, Title: "b", Priority: store.High}})
	if err != nil {
		t.Fatal(err)
	}
	out := b.String()
	if !strings.Contains(out, "PRIORITY") || !strings.Contains(out, "high") {
		t.Fatalf("unexpected table:\n%s", out)
	}
}

func TestTableMarksPinned(t *testing.T) {
	var b strings.Builder
	err := Table(&b, []store.Task{{ID: 1, Title: "pinned", Pinned: true}, {ID: 2, Title: "plain"}})
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimRight(b.String(), "\n"), "\n")
	if len(lines) != 3 {
		t.Fatalf("unexpected table:\n%s", b.String())
	}
	if !strings.HasSuffix(lines[1], PinMark+" pinned") {
		t.Errorf("pinned row has no pin: %q", lines[1])
	}
	if strings.Contains(lines[0], PinMark) || strings.Contains(lines[2], PinMark) {
		t.Errorf("only the pinned row should carry a pin:\n%s", b.String())
	}
	// The pin lives in the last column, so the columns before it still line up.
	if strings.Index(lines[1], PinMark) != strings.Index(lines[2], "plain") {
		t.Errorf("title column is misaligned:\n%s", b.String())
	}
}
