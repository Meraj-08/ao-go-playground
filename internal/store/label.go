package store

import "strings"

// NormalizeLabel lowercases a label and trims the spaces around it.
func NormalizeLabel(label string) string {
	return strings.ToLower(strings.TrimSpace(label))
}
