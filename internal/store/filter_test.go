package store

import "testing"

func TestSearch(t *testing.T) {
	s := &Store{}
	s.Add("Pay the rent", Low)
	urgent := s.Add("Rent a van", High)
	s.Add("Rent a garage", High)
	if err := s.Done(3); err != nil {
		t.Fatal(err)
	}

	if got := s.Search("rent", Filter{}); len(got) != 3 {
		t.Fatalf("the zero filter must keep every match, got %+v", got)
	}
	got := s.Search("rent", Filter{OpenOnly: true, Min: High})
	if len(got) != 1 || got[0].ID != urgent.ID {
		t.Fatalf("got %+v", got)
	}
}
