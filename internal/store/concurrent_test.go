package store

import (
	"fmt"
	"path/filepath"
	"sync"
	"testing"
)

// Two terminals running `taskr add` at the same moment must both keep their
// task, each with its own ID.
func TestConcurrentAddsKeepEveryTask(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	const n = 50
	var wg sync.WaitGroup
	errs := make(chan error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errs <- addOne(path, fmt.Sprintf("task %d", i))
		}(i)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Tasks) != n {
		t.Fatalf("kept %d of %d tasks", len(s.Tasks), n)
	}
	seen := map[int]bool{}
	for _, task := range s.Tasks {
		if seen[task.ID] {
			t.Fatalf("ID %d is used twice", task.ID)
		}
		seen[task.ID] = true
	}
}

// addOne is one `taskr add`: open, add, save.
func addOne(path, title string) error {
	s, err := Open(path)
	if err != nil {
		return err
	}
	s.Add(title, Normal)
	return s.Save()
}
