// Package format renders tasks for the terminal.
package format

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

// Table writes tasks as aligned columns.
func Table(w io.Writer, tasks []store.Task) error {
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tDONE\tPRIORITY\tTITLE")
	for _, t := range tasks {
		mark := " "
		if t.Done {
			mark = "x"
		}
		fmt.Fprintf(tw, "%d\t[%s]\t%s\t%s\n", t.ID, mark, t.Priority, t.Title)
	}
	return tw.Flush()
}
