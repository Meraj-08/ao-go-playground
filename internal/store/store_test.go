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
	a := s.Add("first", Normal)
	s.Add("second", High)
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

func TestParsePriority(t *testing.T) {
	for in, want := range map[string]Priority{"low": Low, "high": High, "normal": Normal, "": Normal} {
		if got, err := ParsePriority(in); err != nil || got != want {
			t.Errorf("ParsePriority(%q) = %v, %v; want %v", in, got, err, want)
		}
	}
	if _, err := ParsePriority("urgent"); err == nil {
		t.Error("expected an error for an unknown priority")
	}
}

func TestSearchMatch(t *testing.T) {
	s := &Store{}
	s.Add("buy milk", Normal)
	s.Add("walk the dog", Normal)
	got := s.Search("milk")
	if len(got) != 1 || got[0].Title != "buy milk" {
		t.Fatalf("unexpected matches: %+v", got)
	}
}

func TestSearchIgnoresCase(t *testing.T) {
	s := &Store{}
	s.Add("Buy Milk", Normal)
	got := s.Search("milk")
	if len(got) != 1 || got[0].Title != "Buy Milk" {
		t.Fatalf("case-insensitive search failed: %+v", got)
	}
}

func TestSearchNoMatch(t *testing.T) {
	s := &Store{}
	s.Add("buy milk", Normal)
	if got := s.Search("bread"); len(got) != 0 {
		t.Fatalf("expected no matches, got %+v", got)
	}
}

func TestSortedPutsOpenHighFirst(t *testing.T) {
	s := &Store{}
	s.Add("low", Low)
	s.Add("high", High)
	done := s.Add("done high", High)
	if err := s.Done(done.ID); err != nil {
		t.Fatal(err)
	}
	got := s.Sorted()
	if got[0].Title != "high" || got[1].Title != "low" || got[2].Title != "done high" {
		t.Fatalf("unexpected order: %v, %v, %v", got[0].Title, got[1].Title, got[2].Title)
	}
}
