package main

import (
	"os"
	"strings"
	"testing"
)

func TestEditWithoutFlagsIsAnError(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(wd) })
	if err := run([]string{"add", "write docs"}); err != nil {
		t.Fatal(err)
	}
	before, err := os.Stat(file)
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	err = run([]string{"edit", "1"})
	if err == nil {
		t.Fatal("edit with no flags: want an error, got nil")
	}
	for _, flag := range []string{"-title", "-priority"} {
		if !strings.Contains(err.Error(), flag) {
			t.Errorf("error %q does not mention %s", err, flag)
		}
	}

	after, err := os.Stat(file)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) || !after.ModTime().Equal(before.ModTime()) {
		t.Error("edit with no flags rewrote the task list")
	}
}
