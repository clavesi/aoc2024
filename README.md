# Advent of Code 2024

These are my solutions to the [2024 Advent of Code](https://adventofcode.com/2024) written to learn and try out Golang.

## Running

In the main `aoc2024/` folder, run

```shell
go run day[DD]/main.go -part [1/2]
```

where `[DD]` represents the day (with leading zero) and `[1/2]` reprsents the part to run which is either part one or part two.

Note that it trys to run using a `input.txt` file in that day's folder. This is not included with this project and thus you have to download your input file from AoC and put it in the folder.

You can of course change the `input` line to run against the `test.txt` given from the problem, but the test file is usually just one or a few lines and does not cover every potential edge cause if you change any of the code.
