// Package store keeps tasks in a JSON file.
package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"
)

// Task is one item on the list.
type Task struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Done     bool      `json:"done"`
	Priority Priority  `json:"priority"`
	Created  time.Time `json:"created"`
}

// Priority orders tasks; higher comes first.
type Priority int

const (
	Low Priority = iota
	Normal
	High
)

// ParsePriority reads "low", "normal" or "high".
func ParsePriority(s string) (Priority, error) {
	switch s {
	case "low":
		return Low, nil
	case "normal", "":
		return Normal, nil
	case "high":
		return High, nil
	}
	return Normal, fmt.Errorf("unknown priority %q (want low, normal or high)", s)
}

// String names the priority the way the CLI accepts it.
func (p Priority) String() string {
	switch p {
	case Low:
		return "low"
	case High:
		return "high"
	}
	return "normal"
}

// Store reads and writes tasks at Path.
type Store struct {
	Path  string
	Tasks []Task
	lock  *os.File
}

// Open loads the tasks at path. A missing file is an empty list.
//
// Open takes the list's lock and Save lets it go, so a command's read, change
// and write happen as one step: a second taskr waits instead of writing over
// the first one's task. Close lets go without saving.
func Open(path string) (*Store, error) {
	lock, err := os.OpenFile(path+".lock", os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, err
	}
	if err := lockFile(lock); err != nil {
		lock.Close()
		return nil, fmt.Errorf("lock %s: %w", path, err)
	}
	s := &Store{Path: path, lock: lock}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		s.Close()
		return nil, err
	}
	if err := json.Unmarshal(data, &s.Tasks); err != nil {
		s.Close()
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return s, nil
}

// Close lets go of the list's lock. Closing a saved store does nothing.
func (s *Store) Close() error {
	if s.lock == nil {
		return nil
	}
	err := s.lock.Close()
	s.lock = nil
	return err
}

// Add appends a task and returns it.
func (s *Store) Add(title string, p Priority) Task {
	t := Task{ID: s.nextID(), Title: title, Priority: p, Created: time.Now()}
	s.Tasks = append(s.Tasks, t)
	return t
}

// Done marks the task with id as done.
func (s *Store) Done(id int) error {
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			s.Tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("no task %d", id)
}

// Sorted returns the tasks with open ones first, then by priority, then oldest first.
func (s *Store) Sorted() []Task {
	out := append([]Task(nil), s.Tasks...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Done != out[j].Done {
			return !out[i].Done
		}
		if out[i].Priority != out[j].Priority {
			return out[i].Priority > out[j].Priority
		}
		return out[i].Created.Before(out[j].Created)
	})
	return out
}

// Save writes the tasks back to disk and lets go of the lock.
func (s *Store) Save() error {
	defer s.Close()
	data, err := json.MarshalIndent(s.Tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0o644)
}

func (s *Store) nextID() int {
	max := 0
	for _, t := range s.Tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}
