package main

import (
	"os"
	"testing"

	"github.com/Meraj-08/ao-go-playground/internal/store"
)

func TestPinUnpin(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(wd) })

	if err := run([]string{"add", "first"}); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		cmd  string
		want bool
	}{{"pin", true}, {"unpin", false}} {
		if err := run([]string{tc.cmd, "1"}); err != nil {
			t.Fatal(err)
		}
		s, err := store.Open(file)
		if err != nil {
			t.Fatal(err)
		}
		if s.Tasks[0].Pinned != tc.want {
			t.Fatalf("after %s: pinned = %v", tc.cmd, s.Tasks[0].Pinned)
		}
		for _, bad := range [][]string{{tc.cmd}, {tc.cmd, "x"}, {tc.cmd, "9"}} {
			if err := run(bad); err == nil {
				t.Errorf("run(%q): expected an error", bad)
			}
		}
	}
}
