package format

import "time"

// Due renders a due date for the list view.
func Due(date time.Time) string {
  if date.IsZero() {
    return ""
  }
  return date.Format("Jan 2")
}
