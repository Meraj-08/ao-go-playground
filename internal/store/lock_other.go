//go:build !unix

package store

import "os"

// lockFile takes no lock where flock doesn't exist, so two adds can still race there.
func lockFile(*os.File) error { return nil }
