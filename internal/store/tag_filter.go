package store

// HasTag reports whether tags holds tag, ignoring case and spacing.
func HasTag(tags []string, tag string) bool {
	want := NormalizeTag(tag)
	for _, have := range tags {
		if NormalizeTag(have) == want {
			return true
		}
	}
	return false
}
