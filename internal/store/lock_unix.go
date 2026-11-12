//go:build unix

package store

import (
	"os"
	"syscall"
)

// lockFile waits until no other taskr holds the list's lock, then takes it.
// The lock goes when the file closes, so a crashed taskr never leaves it held.
func lockFile(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
}
