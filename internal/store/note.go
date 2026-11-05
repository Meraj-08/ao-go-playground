package store

import "strings"

// CleanNote trims a task note and collapses its blank lines.
func CleanNote(note string) string {
	lines := strings.Split(strings.TrimSpace(note), "\n")
	out := lines[:0]
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			out = append(out, strings.TrimRight(line, " "))
		}
	}
	return strings.Join(out, "\n")
}
