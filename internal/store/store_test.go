package store

import (
	"path/filepath"
	"testing"
)

func TestAddDoneSave(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	a := s.Add("first")
	s.Add("second")
	if a.ID != 1 || len(s.Tasks) != 2 {
		t.Fatalf("got id %d and %d tasks", a.ID, len(s.Tasks))
	}
	if err := s.Done(1); err != nil {
		t.Fatal(err)
	}
	if err := s.Save(); err != nil {
		t.Fatal(err)
	}
	again, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if !again.Tasks[0].Done || again.Tasks[1].Done {
		t.Fatalf("done flags not kept: %+v", again.Tasks)
	}
}

func TestDoneMissing(t *testing.T) {
	s := &Store{}
	if err := s.Done(9); err == nil {
		t.Fatal("expected an error for a missing task")
	}
}
