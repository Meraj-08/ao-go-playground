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
