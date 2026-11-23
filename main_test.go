package main

import (
	"io"
	"os"
	"slices"
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

// listOutput runs taskr with args, fails the test if the command errors, and
// returns what it printed to stdout.
//
// For "list" that is the header line followed by one row per task in
// store.Sorted order: pinned first, then open before done, then highest
// priority, then oldest. "list -pinned" drops the unpinned rows and keeps the
// rest in that same order.
func listOutput(t *testing.T, args ...string) string {
	t.Helper()
	stdout, _, err := capture(t, args...)
	if err != nil {
		t.Fatal(err)
	}
	return stdout
}

// capture runs taskr with args and returns what it printed to stdout and to
// stderr, along with the error run returned.
func capture(t *testing.T, args ...string) (stdout, stderr string, runErr error) {
	t.Helper()
	outR, outW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	errR, errW, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	oldOut, oldErr := os.Stdout, os.Stderr
	os.Stdout, os.Stderr = outW, errW
	runErr = run(args)
	os.Stdout, os.Stderr = oldOut, oldErr
	outW.Close()
	errW.Close()
	out, err := io.ReadAll(outR)
	if err != nil {
		t.Fatal(err)
	}
	errOut, err := io.ReadAll(errR)
	if err != nil {
		t.Fatal(err)
	}
	return string(out), string(errOut), runErr
}

// pinnedIDs returns the ids of the pinned tasks on disk, in id order.
func pinnedIDs(t *testing.T) []int {
	t.Helper()
	s, err := store.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	ids := []int{}
	for _, task := range s.Tasks {
		if task.Pinned {
			ids = append(ids, task.ID)
		}
	}
	return ids
}

func TestPinUnpinSeveralIDs(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(wd) })

	for _, title := range []string{"one", "two", "three", "four"} {
		if err := run([]string{"add", title}); err != nil {
			t.Fatal(err)
		}
	}

	for _, tc := range []struct {
		name       string
		args       []string
		wantOut    string
		wantErrOut string
		wantErr    string
		wantPinned []int
	}{
		{
			name:       "pin several, reported in the order given",
			args:       []string{"pin", "4", "1", "2"},
			wantOut:    "pinned #4\npinned #1\npinned #2\n",
			wantPinned: []int{1, 2, 4},
		},
		{
			name:       "unpin several",
			args:       []string{"unpin", "1", "4"},
			wantOut:    "unpinned #1\nunpinned #4\n",
			wantPinned: []int{2},
		},
		{
			name:       "single id still works and is reported",
			args:       []string{"pin", "3"},
			wantOut:    "pinned #3\n",
			wantPinned: []int{2, 3},
		},
		{
			name:       "repeated and already-pinned ids are fine",
			args:       []string{"pin", "3", "3"},
			wantOut:    "pinned #3\npinned #3\n",
			wantPinned: []int{2, 3},
		},
		{
			name:       "bad ids are reported and the good ones still apply",
			args:       []string{"pin", "1", "9", "x", "4"},
			wantOut:    "pinned #1\npinned #4\n",
			wantErrOut: "taskr: no task 9\ntaskr: bad id \"x\"\n",
			wantErr:    "pin failed for 2 of 4 ids",
			wantPinned: []int{1, 2, 3, 4},
		},
		{
			name:       "unpin with a bad id",
			args:       []string{"unpin", "9", "2", "3"},
			wantOut:    "unpinned #2\nunpinned #3\n",
			wantErrOut: "taskr: no task 9\n",
			wantErr:    "unpin failed for 1 of 3 ids",
			wantPinned: []int{1, 4},
		},
		{
			name:       "every id bad changes nothing",
			args:       []string{"unpin", "9", "1.5"},
			wantErrOut: "taskr: no task 9\ntaskr: bad id \"1.5\"\n",
			wantErr:    "unpin failed for 2 of 2 ids",
			wantPinned: []int{1, 4},
		},
		{
			name:       "no ids",
			args:       []string{"pin"},
			wantErr:    "pin needs at least one id",
			wantPinned: []int{1, 4},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out, errOut, err := capture(t, tc.args...)
			if out != tc.wantOut {
				t.Errorf("stdout = %q, want %q", out, tc.wantOut)
			}
			if errOut != tc.wantErrOut {
				t.Errorf("stderr = %q, want %q", errOut, tc.wantErrOut)
			}
			switch {
			case tc.wantErr == "" && err != nil:
				t.Errorf("unexpected error: %v", err)
			case tc.wantErr != "" && (err == nil || err.Error() != tc.wantErr):
				t.Errorf("error = %v, want %q", err, tc.wantErr)
			}
			if got := pinnedIDs(t); !slices.Equal(got, tc.wantPinned) {
				t.Errorf("pinned ids = %v, want %v", got, tc.wantPinned)
			}
		})
	}
}

// With nothing to save, a failed pin must not create tasks.json.
func TestPinAllBadIDsDoesNotWriteFile(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(wd) })

	if _, _, err := capture(t, "pin", "1", "2"); err == nil {
		t.Fatal("expected an error")
	}
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		t.Errorf("tasks.json should not exist, stat err = %v", err)
	}
}

func TestRunWithoutArgs(t *testing.T) {
	if err := run(nil); err == nil {
		t.Fatal("expected a usage error")
	}
}
