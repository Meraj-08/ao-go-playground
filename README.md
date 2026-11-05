# taskr

A tiny todo CLI used to try out the Agent Orchestrator PR view.

## Usage

| Command | What it does |
| --- | --- |
| `taskr add "<title>"` | Adds a task with normal priority |
| `taskr add -priority high "<title>"` | Adds a task with the given priority |
| `taskr list` | Shows all tasks as a table, open tasks first, highest priority at the top |
| `taskr find <words>` | Shows the tasks whose titles contain every word, in the same order as `taskr list` |
| `taskr find -open <words>` | Leaves out tasks that are done |
| `taskr find -min high <words>` | Leaves out tasks below the given priority (`low` by default, so nothing is left out) |
| `taskr done <id>` | Marks a task as done |

Priorities are `low`, `normal` (the default) and `high`; anything else is rejected.

`taskr find` ignores case and matches parts of words, so `rent` also finds "Parent evening". If nothing matches, it prints `no tasks match "<words>"`. Put flags before the words: in `taskr find rent -open`, `-open` is searched for as text.

Tasks are kept in `tasks.json` in the current directory.

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
