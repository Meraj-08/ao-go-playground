package format

import (
    "fmt"
    "io"

    "github.com/Meraj-08/ao-go-playground/internal/store"
)

// Change writes what an edit changed, one line per field, or "no change".
func Change(w io.Writer, before, after store.Task) error {
    changed := false
    if before.Title != after.Title {
        fmt.Fprintf(w, "#%d title: %q -> %q\n", after.ID, before.Title, after.Title)
        changed = true
    }
    if before.Priority != after.Priority {
        fmt.Fprintf(w, "#%d priority: %s -> %s\n", after.ID, before.Priority, after.Priority)
        changed = true
    }
    if !changed {
        _, err := fmt.Fprintf(w, "#%d: no change\n", after.ID)
        return err
    }
    return nil
}
