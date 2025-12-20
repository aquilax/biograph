# biograph [![Go Reference](https://pkg.go.dev/badge/github.com/aquilax/biograph.svg)](https://pkg.go.dev/github.com/aquilax/biograph)

Life graphing library

## Reporters

- `Branches` — ASCII graph showing event branches and merges by start/end dates. Create with `report.NewBranches(start, end, io.Writer)` and call `Generate(events)` to produce the timeline.

