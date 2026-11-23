package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestMain lets the test binary act as taskr, so tests can check real output.
func TestMain(m *testing.M) {
	if os.Getenv("TASKR_RUN_MAIN") == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestCountCommand(t *testing.T) {
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ name, tasks, want string }{
		{"no file", "", "0 open, 0 done\n"},
		{"mixed", `[{"id":1},{"id":2,"done":true},{"id":3}]`, "2 open, 1 done\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			if tc.tasks != "" {
				if err := os.WriteFile(filepath.Join(dir, file), []byte(tc.tasks), 0o644); err != nil {
					t.Fatal(err)
				}
			}
			cmd := exec.Command(exe, "count")
			cmd.Dir = dir
			cmd.Env = append(os.Environ(), "TASKR_RUN_MAIN=1")
			out, err := cmd.Output()
			if err != nil {
				t.Fatal(err)
			}
			if string(out) != tc.want {
				t.Fatalf("got %q, want %q", out, tc.want)
			}
		})
	}
}
