package store

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestMain doubles as the process that fails: with TASKR_SAVE_PATH set, it adds
// a task and saves, so a test can run a real save with no room to write.
func TestMain(m *testing.M) {
	if path := os.Getenv("TASKR_SAVE_PATH"); path != "" {
		s, err := Open(path)
		if err == nil {
			s.Add("fourth", Normal)
			err = s.Save()
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	os.Exit(m.Run())
}

// A file size limit of zero fails every write, the way a full disk does.
func TestFailedSaveKeepsTasks(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, title := range []string{"first", "second", "third"} {
		s.Add(title, Normal)
	}
	if err := s.Save(); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("sh", "-c", `ulimit -f 0; exec "$0"`, os.Args[0])
	cmd.Env = append(os.Environ(), "TASKR_SAVE_PATH="+path)
	if err := cmd.Run(); err == nil {
		t.Fatal("the save should have failed")
	}

	again, err := Open(path)
	if err != nil {
		t.Fatalf("reopen after a failed save: %v", err)
	}
	if len(again.Tasks) != 3 {
		t.Fatalf("tasks.json has %d tasks after a failed save, want 3", len(again.Tasks))
	}
}
