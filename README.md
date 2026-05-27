# learning-go-i-guess
ok so im learning go.

im a self taught python dev and go looks weird as hell but fast.
turns out go is actually super fun once you get past the pointers and implicit interfaces.

## project structure
- `basics/`: variables, types, constants, and console input
- `control/`: loops, conditionals, switch
- `functions/`: standard funcs, variadic, closures, recursion
- `data_structures/`: arrays, slices, maps, structs
- `pointers/`: pointers_wtf.go, pointer_funcs.go
- `interfaces/`: interfaces, stringer, type assertion
- `errors/`: error handling, custom errors, panic/recover
- `concurrency/`: goroutines, waitgroups, channels, select, mutex
- `packages/`: basic custom packages
- `stdlib_exploration/`: JSON, HTTP servers, regexp, times
- `mini_projects/`:
  - `todo_cli/`: JSON-persisted tasks list
  - `expense_tracker/`: CLI for tracking spending by category
  - `url_shortener/`: in-memory URL shortener HTTP API
- `testing/`: unit tests and table-driven tests

## lessons learned
1. `:=` is cool, but compiler will yell at you if you don't use it.
2. Slices are references. If you mutate a sliced slice, you mutate the original. Cursed.
3. Interfaces are implicit. No `implements` keyword. Just write the methods and it works.
4. Error handling: `if err != nil` is 80% of Go code. Get used to it.
5. Concurrency in Go is cheat code. Waitgroups and Channels are awesome.

run `go run main.go` to see the interactive launcher.
