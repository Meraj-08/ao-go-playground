package store

import "time"

// ArchiveAfter is how long a finished task stays in the main list.
const ArchiveAfter = 30 * 24 * time.Hour

// ShouldArchive reports whether a finished task is old enough to archive.
func ShouldArchive(done bool, finishedAt, now time.Time) bool {
	return done && now.Sub(finishedAt) > ArchiveAfter
}
