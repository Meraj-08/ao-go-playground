package store

// Filter narrows a search. The zero value keeps every match.
type Filter struct {
	OpenOnly bool
	Min      Priority
}

// Search returns the tasks matching query that also pass f.
func (s *Store) Search(query string, f Filter) []Task {
	var out []Task
	for _, t := range s.Find(query) {
		if f.OpenOnly && t.Done {
			continue
		}
		if t.Priority < f.Min {
			continue
		}
		out = append(out, t)
	}
	return out
}
