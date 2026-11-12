package store

import "testing"

func TestRename(t *testing.T) {
	s := &Store{}
	s.Add("Pay rent", Normal)
	if err := s.Rename(1, "Pay the rent"); err != nil {
		t.Fatal(err)
	}
	if s.Tasks[0].Title != "Pay the rent" {
		t.Fatalf("title = %q", s.Tasks[0].Title)
	}
	if err := s.Rename(1, "  "); err == nil {
		t.Fatal("a blank title must be refused")
	}
	if err := s.Rename(9, "Nope"); err == nil {
		t.Fatal("an unknown id must be refused")
	}
}

func TestSetPriority(t *testing.T) {
	s := &Store{}
	s.Add("Book dentist", Normal)
	s.Add("Renew passport", Normal)
	if err := s.SetPriority(2, High); err != nil {
		t.Fatal(err)
	}
	if got := s.Sorted()[0].Title; got != "Renew passport" {
		t.Fatalf("first after raising #2 = %q, want Renew passport", got)
	}
	if err := s.SetPriority(9, Low); err == nil {
		t.Fatal("an unknown id must be refused")
	}
}
