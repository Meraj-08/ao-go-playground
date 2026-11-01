package format

import (
	"strings"
	"testing"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

func TestTable(t *testing.T) {
	var b strings.Builder
	err := Table(&b, []store.Task{{ID: 1, Title: "a", Done: true}, {ID: 2, Title: "b"}})
	if err != nil {
		t.Fatal(err)
	}
	out := b.String()
	if !strings.Contains(out, "[x]   a") || !strings.Contains(out, "[ ]   b") {
		t.Fatalf("unexpected table:\n%s", out)
	}
}
