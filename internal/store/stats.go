package store

// Stats counts tasks by state.
type Stats struct {
	Total, Open, Done int
}

// Stats returns how many tasks are open and done.
func (s *Store) Stats() Stats {
	st := Stats{Total: len(s.Tasks)}
	for _, t := range s.Tasks {
		if t.Done {
			st.Done++
		} else {
			st.Open++
		}
	}
	return st
}
