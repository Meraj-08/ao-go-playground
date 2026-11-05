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

## Task notes

A task note is cleaned up by trimming surrounding spaces and dropping blank lines, and can be at most 280 characters; anything longer is cut off. Notes display like `note: call the plumber`, and an empty note displays as blank.

This is handled by helpers in `internal/store` and `internal/format` for now: the `taskr` commands don't accept or show notes yet.

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
