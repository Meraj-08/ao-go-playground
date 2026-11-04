package format

import "strings"

// Tags renders a task's tags for the list view.
func Tags(tags []string) string {
  if len(tags) == 0 {
    return ""
  }
  return "#" + strings.Join(tags, " #")
}
