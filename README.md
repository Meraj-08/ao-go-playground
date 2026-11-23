# taskr

A tiny todo CLI used to try out the Agent Orchestrator PR view.

## Usage

| Command | What it does |
| --- | --- |
| `taskr add "<title>"` | Adds a task with normal priority |
| `taskr add -priority high "<title>"` | Adds a task with the given priority |
| `taskr list` | Shows all tasks as a table: pinned tasks first, then open before done, then highest priority |
| `taskr list -pinned` | Shows only the pinned tasks, in the same table and order |
| `taskr done <id>` | Marks a task as done |
| `taskr pin <id>...` | Pins one or more tasks so they are listed first |
| `taskr unpin <id>...` | Unpins one or more tasks |

Priorities are `low`, `normal` (the default) and `high`; anything else is rejected.

Pinned tasks are marked with a 📌 before the title and always sit above unpinned ones, even when they are done. Within the pinned and unpinned groups the usual order applies: open before done, then by priority, then oldest first.

```
ID  DONE  PRIORITY  TITLE
3   [ ]   low       📌 water the plants
1   [ ]   high      ship the release
2   [x]   normal    write the changelog
```

`taskr list -pinned` leaves out the unpinned rows and changes nothing else; with no pinned tasks it prints just the header.

```
ID  DONE  PRIORITY  TITLE
3   [ ]   low       📌 water the plants
```

`pin` and `unpin` take any number of ids and report each one on its own line, in the order given:

```
$ taskr pin 2 5 7
pinned #2
pinned #5
pinned #7
```

An id that is not a number or matches no task is reported on stderr and does not stop the others: the rest are still pinned (or unpinned) and saved, and `taskr` exits with status 1.

```
$ taskr pin 2 9 7
pinned #2
taskr: no task 9
pinned #7
taskr: pin failed for 1 of 3 ids
```

Tasks are kept in `tasks.json` in the current directory.

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
