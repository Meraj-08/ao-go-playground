# taskr

A tiny todo CLI used to try out the Agent Orchestrator PR view.

## Usage

| Command | What it does |
| --- | --- |
| `taskr add "<title>"` | Adds a task with normal priority |
| `taskr add -priority high "<title>"` | Adds a task with the given priority |
| `taskr list` | Shows all tasks as a table, open tasks first, highest priority at the top |
| `taskr done <id>` | Marks a task as done |

Priorities are `low`, `normal` (the default) and `high`; anything else is rejected.

Tasks are kept in `tasks.json` in the current directory.

## Task labels

Labels are lowercased and have surrounding spaces trimmed, so ` Work ` and `work` are the same label. A set of labels is sorted alphabetically with duplicates and empty labels removed, and displays like `[home, work]`; no labels displays as blank.

This is handled by helpers in `internal/store` and `internal/format` for now: the `taskr` commands don't accept or show labels yet.

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
