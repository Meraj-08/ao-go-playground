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

	out := listOutput(t, "list")

	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
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

func TestListPinnedFlag(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(wd) })

	const header = "ID  DONE  PRIORITY  TITLE"
	if got := listOutput(t, "list", "-pinned"); got != header+"\n" {
		t.Errorf("empty store: got %q, want only the header", got)
	}

	for _, args := range [][]string{
		{"add", "-priority", "high", "urgent"},
		{"add", "-priority", "low", "pinned low"},
		{"add", "-priority", "high", "pinned high"},
		{"add", "pinned done"},
		{"pin", "2"},
		{"pin", "3"},
		{"pin", "4"},
		{"done", "4"},
	} {
		if err := run(args); err != nil {
			t.Fatal(err)
		}
	}

	// No pinned filter: all four rows. With it: the same rows minus the
	// unpinned one, in the same order.
	all := strings.Split(strings.TrimRight(listOutput(t, "list"), "\n"), "\n")
	if len(all) != 5 {
		t.Fatalf("unexpected list output: %q", all)
	}
	var want []string
	for _, l := range all {
		if l == header || strings.Contains(l, "📌") {
			want = append(want, l)
		}
	}
	if len(want) != 4 {
		t.Fatalf("expected three pinned rows in %q", all)
	}
	got := strings.Split(strings.TrimRight(listOutput(t, "list", "-pinned"), "\n"), "\n")
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("list -pinned:\n%s\nwant:\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
	for i, title := range []string{"pinned high", "pinned low", "pinned done"} {
		if !strings.HasSuffix(got[i+1], "📌 "+title) {
			t.Errorf("row %d = %q, want title %q", i+1, got[i+1], title)
		}
	}

	for _, bad := range [][]string{
		{"list", "-bogus"},
		{"list", "pinned"},
		{"list", "-pinned", "extra"},
		{"list", "-pinned=maybe"},
	} {
		if err := run(bad); err == nil {
			t.Errorf("run(%q): expected an error", bad)
		}
	}
}

// listOutput runs taskr with args and returns what it printed to stdout.
func listOutput(t *testing.T, args ...string) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w
	runErr := run(args)
	os.Stdout = stdout
	w.Close()
	if runErr != nil {
		t.Fatal(runErr)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func TestRunWithoutArgs(t *testing.T) {
	if err := run(nil); err == nil {
		t.Fatal("expected a usage error")
	}
}
