# taskr

A tiny todo CLI used to try out the Agent Orchestrator PR view.

## Usage

| Command | What it does |
| --- | --- |
| `taskr add "<title>"` | Adds a task with normal priority |
| `taskr add -priority high "<title>"` | Adds a task with the given priority |
| `taskr list` | Shows all tasks as a table, open tasks first, highest priority at the top |
| `taskr done <id>` | Marks a task as done |
| `taskr edit <id> -title "<title>"` | Renames a task |
| `taskr edit <id> -priority <priority>` | Changes a task's priority |

Priorities are `low`, `normal` (the default) and `high`; anything else is rejected.

`taskr edit` takes the id first, then `-title`, `-priority` or both, and prints one line for each field it changed (`#1 priority: normal -> high`), or `#1: no change` when nothing differs.

Tasks are kept in `tasks.json` in the current directory.

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
