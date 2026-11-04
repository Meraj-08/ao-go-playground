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

## Tags

Tags are not part of the CLI yet: tasks don't store them, no command sets them, and `taskr list` doesn't show them. So far they exist as helpers in `internal/`:

| Helper | What it does |
| --- | --- |
| `store.NormalizeTag` | Lowercases a tag and trims the space around it |
| `store.HasTag` | Reports whether a list holds a tag, ignoring case and surrounding space |
| `format.Tags` | Renders tags as `#work #home`, and returns an error if any tag is blank |
| `format.TagCount` | Renders the number of tags, such as `2 tags` |

## Development

```sh
make test   # go vet + go test
make build  # builds ./taskr
```
