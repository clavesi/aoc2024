# Advent of Code 2024

These are my solutions to the [2024 Advent of Code](https://adventofcode.com/2024) written to learn and try out Golang.

## Running

### Code

In the main `aoc2024/` folder, run

```shell
go run day[DD]/main.go -part=[1/2]
```

where `[DD]` represents the day (with leading zero) and `[1/2]` reprsents the part to run which is either part one or part two.

Note that it trys to run using a `input.txt` file in that day's folder. This is not included with this project and thus you have to download your own input file from the AoC website and put it in the folder.

You can of course change the `input` line to run against the `test.txt` given from the problem, but the test file is usually just one or a few lines and does not cover every potential edge cause if you change any of the code.

---

### Test

In the main `aoc2024/` directory

To run an **entire day's** tests,

```shell
go test -v ./day[DD]
```

where `[DD]` represents the day (with leading zero). `-v` just indicates verbose, giving some more detail to the output and can be omitted if desired.

To run an **individual** test,

```shell
go test -v ./day[DD] -run [TEST_NAME]
```

where `[TEST_NAME]` is the name of the function for the test you want to run. In general, each day will have at least two tests cases, one for each part with the given example.

To run **every** test,

```shell
go test -v ./...
```

Take care to note that it's three dots after the forward slash.
