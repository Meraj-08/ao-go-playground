package format

import (
	"strings"
	"testing"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

func TestChange(t *testing.T) {
	var b strings.Builder
	before := store.Task{ID: 2, Title: "Pay rent", Priority: store.Normal}
	after := store.Task{ID: 2, Title: "Pay the rent", Priority: store.High}
	if err := Change(&b, before, after); err != nil {
		t.Fatal(err)
	}
	want := "#2 title: \"Pay rent\" -> \"Pay the rent\"\n#2 priority: normal -> high\n"
	if b.String() != want {
		t.Fatalf("got %q, want %q", b.String(), want)
	}
}
