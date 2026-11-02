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

func TestDelete(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	s.Add("first")
	s.Add("second")
	s.Add("third")
	if err := s.Delete(2); err != nil {
		t.Fatal(err)
	}
	if err := s.Save(); err != nil {
		t.Fatal(err)
	}
	again, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(again.Tasks) != 2 || again.Tasks[0].ID != 1 || again.Tasks[1].ID != 3 {
		t.Fatalf("wrong tasks after delete: %+v", again.Tasks)
	}
	if next := again.Add("fourth"); next.ID != 4 {
		t.Fatalf("got id %d after delete, want 4", next.ID)
	}
}

func TestDeleteMissing(t *testing.T) {
	s := &Store{}
	s.Add("only")
	for _, id := range []int{9, 0, -1} {
		if err := s.Delete(id); err == nil {
			t.Fatalf("expected an error deleting task %d", id)
		}
	}
	if len(s.Tasks) != 1 {
		t.Fatalf("tasks changed after failed deletes: %+v", s.Tasks)
	}
}

func TestDoneMissing(t *testing.T) {
	s := &Store{}
	if err := s.Done(9); err == nil {
		t.Fatal("expected an error for a missing task")
	}
}
