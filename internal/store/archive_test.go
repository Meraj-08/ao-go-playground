package store

import (
	"testing"
	"time"
)

func TestShouldArchive(t *testing.T) {
	now := time.Date(2026, 11, 5, 0, 0, 0, 0, time.UTC)
	if !ShouldArchive(true, now.Add(-31*24*time.Hour), now) {
		t.Fatal("an old finished task should be archived")
	}
	if ShouldArchive(true, now.Add(-time.Hour), now) || ShouldArchive(false, now.Add(-90*24*time.Hour), now) {
		t.Fatal("recent or open tasks should stay")
	}
}
