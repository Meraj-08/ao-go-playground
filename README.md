# taskr

A tiny todo CLI used to try out the Agent Orchestrator PR view.

```sh
go run . add "write the docs"
go run . list
go run . done 1
go run . export > tasks.csv
```

Tasks are kept in `tasks.json` in the current directory.
