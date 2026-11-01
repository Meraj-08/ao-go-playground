# taskr

A tiny todo CLI used to try out the Agent Orchestrator PR view.

## Usage

| Command | What it does |
| --- | --- |
| `taskr add "<title>"` | Adds a task |
| `taskr list` | Shows all tasks as a table |
| `taskr done <id>` | Marks a task as done |

Tasks are kept in `tasks.json` in the current directory.

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
