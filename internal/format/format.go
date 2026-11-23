// Package format renders tasks for the terminal.
package format

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

// PinMark is shown in front of the title of a pinned task.
const PinMark = "📌"

// Table writes tasks as aligned columns, one row per task in the order given;
// it does not sort, so pass store.Sorted() for the usual pinned-first order.
// Pinned tasks get PinMark before the title; it sits in the last column so its
// width cannot skew the alignment.
func Table(w io.Writer, tasks []store.Task) error {
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	fmt.Fprintln(tw, "ID\tDONE\tPRIORITY\tTITLE")
	for _, t := range tasks {
		mark := " "
		if t.Done {
			mark = "x"
		}
		title := t.Title
		if t.Pinned {
			title = PinMark + " " + title
		}
		fmt.Fprintf(tw, "%d\t[%s]\t%s\t%s\n", t.ID, mark, t.Priority, title)
	}
	return tw.Flush()
}
