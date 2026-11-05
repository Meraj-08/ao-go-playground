package format

// Note prefixes a task note for the task list.
func Note(note string) string {
  if note == "" {
    return ""
  }
  return "note: " + note
}
