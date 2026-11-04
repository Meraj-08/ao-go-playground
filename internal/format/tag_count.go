package format

import "fmt"

// TagCount renders how many tags a task has.
func TagCount(tags []string) string {
  return fmt.Sprintf("%d tags", len(tags))
}
