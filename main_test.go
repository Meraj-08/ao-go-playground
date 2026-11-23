package main

import (
	"os"
	"strings"
	"testing"
)

func TestRunUndo(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(wd)

	// Steps share one tasks.json; an empty want means success.
	for _, step := range []struct{ args, want string }{
		{"", "usage:"},
		{"add task", ""},
		{"undo 1", "task 1 is already open"},
		{"done 1", ""},
		{"undo 1", ""},
		{"undo 1", "task 1 is already open"},
		{"undo 9", "no task 9"},
		{"undo", "undo needs an id"},
		{"undo x", `bad id "x"`},
	} {
		err := run(strings.Fields(step.args))
		if step.want == "" && err != nil {
			t.Fatalf("run(%q): %v", step.args, err)
		}
		if step.want != "" && (err == nil || !strings.Contains(err.Error(), step.want)) {
			t.Fatalf("run(%q) = %v; want %q", step.args, err, step.want)
		}
	}
}
