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
| `taskr pin <id>` | Pins a task so it is listed first |
| `taskr unpin <id>` | Unpins a task |

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

Tasks are kept in `tasks.json` in the current directory.

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
