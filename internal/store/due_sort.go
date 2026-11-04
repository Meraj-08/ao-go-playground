package store

import (
	"sort"
	"time"
)

// SortDue orders dates with the soonest first.
func SortDue(dates []time.Time) {
	sort.Slice(dates, func(i, j int) bool { return dates[i].Before(dates[j]) })
}
