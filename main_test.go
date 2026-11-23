package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

func TestPinUnpin(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(wd) })

	if err := run([]string{"add", "first"}); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		cmd  string
		want bool
	}{{"pin", true}, {"unpin", false}} {
		if err := run([]string{tc.cmd, "1"}); err != nil {
			t.Fatal(err)
		}
		s, err := store.Open(file)
		if err != nil {
			t.Fatal(err)
		}
		if s.Tasks[0].Pinned != tc.want {
			t.Fatalf("after %s: pinned = %v", tc.cmd, s.Tasks[0].Pinned)
		}
		for _, bad := range [][]string{{tc.cmd}, {tc.cmd, "x"}, {tc.cmd, "9"}} {
			if err := run(bad); err == nil {
				t.Errorf("run(%q): expected an error", bad)
			}
		}
	}
}

func TestListShowsPinnedFirst(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(wd) })

	for _, args := range [][]string{
		{"add", "-priority", "high", "urgent"},
		{"add", "-priority", "low", "keep in view"},
		{"pin", "2"},
	} {
		if err := run(args); err != nil {
			t.Fatal(err)
		}
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w
	runErr := run([]string{"list"})
	os.Stdout = stdout
	w.Close()
	if runErr != nil {
		t.Fatal(runErr)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(strings.TrimRight(string(out), "\n"), "\n")
	if len(lines) != 3 {
		t.Fatalf("unexpected list output:\n%s", out)
	}
	if !strings.HasSuffix(lines[1], "📌 keep in view") {
		t.Errorf("pinned task should be first and marked, got %q", lines[1])
	}
	if !strings.HasSuffix(lines[2], "urgent") || strings.Contains(lines[2], "📌") {
		t.Errorf("unpinned task should follow without a pin, got %q", lines[2])
	}
}

func TestRunWithoutArgs(t *testing.T) {
	if err := run(nil); err == nil {
		t.Fatal("expected a usage error")
	}
}
