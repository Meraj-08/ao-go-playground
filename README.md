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

## Due dates

Due dates are written as `YYYY-MM-DD`, such as `2026-11-05`; surrounding spaces are ignored and anything else is rejected. They sort soonest first and display like `Nov 5`, and an unset date displays as blank.

This is handled by helpers in `internal/store` and `internal/format` for now: the `taskr` commands don't accept or show due dates yet.

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
