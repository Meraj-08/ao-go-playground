package store

import (
	"fmt"
	"strings"
)

// Rename gives the task with id a new title. A blank title is refused, since
// a task with no words can't be told apart in the list.
func (s *Store) Rename(id int, title string) error {
	if strings.TrimSpace(title) == "" {
		return fmt.Errorf("a task needs a title")
	}
	t, err := s.find(id)
	if err != nil {
		return err
	}
	t.Title = title
	return nil
}

// SetPriority moves the task with id to priority p, which reorders the list.
func (s *Store) SetPriority(id int, p Priority) error {
	t, err := s.find(id)
	if err != nil {
		return err
	}
	t.Priority = p
	return nil
}

// find returns the task with id, to change in place.
func (s *Store) find(id int) (*Task, error) {
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			return &s.Tasks[i], nil
		}
	}
	return nil, fmt.Errorf("no task %d", id)
}
