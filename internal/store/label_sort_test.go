package store

import (
	"reflect"
	"testing"
)

func TestSortLabels(t *testing.T) {
	got := SortLabels([]string{"Work", " home", "work"})
	if want := []string{"home", "work"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v", got)
	}
}
