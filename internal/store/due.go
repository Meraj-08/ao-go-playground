package store

import (
	"strings"
	"time"
)

// ParseDue reads a due date written as YYYY-MM-DD.
func ParseDue(value string) (time.Time, error) {
	return time.Parse(time.DateOnly, strings.TrimSpace(value))
}
