package store

// MaxNote is the longest note a task can keep.
const MaxNote = 280

// LimitNote cuts a note down to MaxNote characters.
func LimitNote(note string) string {
	runes := []rune(note)
	if len(runes) <= MaxNote {
		return note
	}
	return string(runes[:MaxNote])
}
