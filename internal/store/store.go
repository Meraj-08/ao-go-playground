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

// ParsePriority reads "low", "normal" or "high". Anything else is Normal.
func ParsePriority(s string) Priority {
	switch s {
	case "low":
		return Low
	case "high":
		return High
	}
	return Normal
}

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
}

// Open loads the tasks at path. A missing file is an empty list.
func Open(path string) (*Store, error) {
	s := &Store{Path: path}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &s.Tasks); err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return s, nil
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

// Save writes the tasks back to disk.
func (s *Store) Save() error {
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
