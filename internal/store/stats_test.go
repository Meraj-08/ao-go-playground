package store

import "testing"

func TestStats(t *testing.T) {
	s := &Store{}
	s.Add("a")
	s.Add("b")
	if err := s.Done(2); err != nil {
		t.Fatal(err)
	}
	if got := s.Stats(); got != (Stats{Total: 2, Open: 1, Done: 1}) {
		t.Fatalf("got %+v", got)
	}
}

func TestStatsEmpty(t *testing.T) {
	if got := (&Store{}).Stats(); got != (Stats{}) {
		t.Fatalf("got %+v", got)
	}
}
