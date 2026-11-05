package store

import "testing"

func TestFind(t *testing.T) {
	s := &Store{}
	s.Add("Pay the rent", Normal)
	s.Add("Rent a van", High)
	s.Add("Book dentist", Normal)

	got := s.Find("RENT")
	if len(got) != 2 || got[0].Title != "Rent a van" {
		t.Fatalf("got %+v", got)
	}
	if got := s.Find("rent van"); len(got) != 1 {
		t.Fatalf("every word must match, got %+v", got)
	}
	if got := s.Find("plumber"); len(got) != 0 {
		t.Fatalf("got %+v", got)
	}
}
