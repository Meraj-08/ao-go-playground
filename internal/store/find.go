package store

import "strings"

// Find returns the tasks whose titles contain every word of query, ignoring
// case, in the same order as Sorted.
func (s *Store) Find(query string) []Task {
	words := strings.Fields(strings.ToLower(query))
	var out []Task
	for _, t := range s.Sorted() {
		if matchesAll(strings.ToLower(t.Title), words) {
			out = append(out, t)
		}
	}
	return out
}

func matchesAll(title string, words []string) bool {
	for _, word := range words {
		if !strings.Contains(title, word) {
			return false
		}
	}
	return true
}
