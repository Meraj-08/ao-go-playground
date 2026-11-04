package store

import "strings"

// NormalizeTag lowercases a tag and trims the space around it.
func NormalizeTag(tag string) string {
	return strings.ToLower(strings.TrimSpace(tag))
}
