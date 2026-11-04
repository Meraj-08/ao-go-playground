package format

import (
	"fmt"
	"strings"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

// Tags renders a task's tags for the list view. It returns an error if any
// tag is blank, since that would otherwise show up as a stray "#".
func Tags(tags []string) (string, error) {
	if len(tags) == 0 {
		return "", nil
	}
	for i, tag := range tags {
		if store.NormalizeTag(tag) == "" {
			return "", fmt.Errorf("tag %d is blank", i+1)
		}
	}
	return "#" + strings.Join(tags, " #"), nil
}
