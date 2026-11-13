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
		return fmt.Errorf("usage: taskr add [-priority p] <title> | list | done <id> | edit <id> [-title t] [-priority p]")
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
		p, err := store.ParsePriority(*prio)
		if err != nil {
			return err
		}
		t := s.Add(fs.Arg(0), p)
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
	case "edit":
		fs := flag.NewFlagSet("edit", flag.ContinueOnError)
		title := fs.String("title", "", "the new title")
		prio := fs.String("priority", "", "low, normal or high")
		if len(args) < 2 {
			return fmt.Errorf("edit needs an id")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("bad id %q", args[1])
		}
		if err := fs.Parse(args[2:]); err != nil {
			return err
		}
		if *title == "" && *prio == "" {
			return fmt.Errorf("edit needs -title or -priority")
		}
		var before store.Task
		for _, t := range s.Tasks {
			if t.ID == id {
				before = t
			}
		}
		if *title != "" {
			if err := s.Rename(id, *title); err != nil {
				return err
			}
		}
		if *prio != "" {
			p, err := store.ParsePriority(*prio)
			if err != nil {
				return err
			}
			if err := s.SetPriority(id, p); err != nil {
				return err
			}
		}
		for _, t := range s.Tasks {
			if t.ID == id {
				if err := format.Change(os.Stdout, before, t); err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
	return s.Save()
}
