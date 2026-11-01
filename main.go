// Command taskr is a tiny todo list.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Meraj-08/ao-go-playground/internal/format"
	"github.com/Meraj-08/ao-go-playground/internal/store"
)

const file = "tasks.json"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "taskr:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: taskr add [-priority p] <title> | list | done <id>")
	}
	s, err := store.Open(file)
	if err != nil {
		return err
	}
	switch args[0] {
	case "add":
		fs := flag.NewFlagSet("add", flag.ContinueOnError)
		prio := fs.String("priority", "normal", "low, normal or high")
		if err := fs.Parse(args[1:]); err != nil {
			return err
		}
		if fs.NArg() < 1 {
			return fmt.Errorf("add needs a title")
		}
		t := s.Add(fs.Arg(0), store.ParsePriority(*prio))
		fmt.Printf("added #%d\n", t.ID)
	case "list":
		return format.Table(os.Stdout, s.Sorted())
	case "done":
		if len(args) < 2 {
			return fmt.Errorf("done needs an id")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("bad id %q", args[1])
		}
		if err := s.Done(id); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
	return s.Save()
}
