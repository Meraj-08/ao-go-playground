package format

import (
  "fmt"
  "io"

  "github.com/Meraj-08/ao-go-playground/internal/store"
)

// Matches writes search results, or says that nothing matched.
func Matches(w io.Writer, query string, tasks []store.Task) error {
  if len(tasks) == 0 {
    _, err := fmt.Fprintf(w, "no tasks match %q\n", query)
    return err
  }
  return Table(w, tasks)
}
