package format

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

// CSV writes tasks as comma-separated rows with a header.
func CSV(w io.Writer, tasks []store.Task) error {
	cw := csv.NewWriter(w)
	cw.Write([]string{"id", "done", "title"})
	for _, t := range tasks {
		cw.Write([]string{strconv.Itoa(t.ID), strconv.FormatBool(t.Done), t.Title})
	}
	cw.Flush()
	return cw.Error()
}

func csvHeader() []string { return []string{"id","done","title"} }
