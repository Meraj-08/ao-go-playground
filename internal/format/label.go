package format

import "strings"

// Labels joins labels for the task list.
func Labels(labels []string) string {
  if len(labels) == 0 {
    return ""
  }
  return "[" + strings.Join(labels, ", ") + "]"
}
