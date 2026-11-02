# taskr

A tiny todo CLI used to try out the Agent Orchestrator PR view.

```sh
go run . add "write the docs"
go run . add -priority high "fix the login bug"
go run . list
go run . done 1
```

Tasks are kept in `tasks.json` in the current directory.

`list` shows open tasks first, highest priority at the top.

Priorities are `low`, `normal` (the default) and `high`; anything else is rejected.
