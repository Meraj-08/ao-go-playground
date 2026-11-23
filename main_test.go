package main

import (
	"io"
	"os"
	"testing"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

// capture runs the CLI and returns what it printed.
func capture(t *testing.T, args ...string) string {
	t.Helper()
	old := os.Stdout
	defer func() { os.Stdout = old }()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	runErr := run(args)
	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if runErr != nil {
		t.Fatal(runErr)
	}
	return string(out)
}

func TestClearDoneCommand(t *testing.T) {
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(old) })

	capture(t, "add", "keep")
	capture(t, "add", "drop")
	capture(t, "done", "2")
	if got := capture(t, "clear-done"); got != "removed 1 done task(s)\n" {
		t.Fatalf("printed %q", got)
	}
	s, err := store.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Tasks) != 1 || s.Tasks[0].Title != "keep" {
		t.Fatalf("saved tasks: %+v", s.Tasks)
	}
}
