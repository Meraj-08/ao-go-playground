package store

import (
	"testing"
	"time"
)

func TestSortDue(t *testing.T) {
	soon := time.Date(2026, 11, 5, 0, 0, 0, 0, time.UTC)
	later := soon.AddDate(0, 0, 3)
	dates := []time.Time{later, soon}
	SortDue(dates)
	if !dates[0].Equal(soon) {
		t.Fatalf("first = %v, want %v", dates[0], soon)
	}
}
