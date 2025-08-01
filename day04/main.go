package main

import (
	"aoc/utils"
	"flag"
	"fmt"
)

type Direction [2]int

var (
	North     = Direction{0, 1}
	East      = Direction{1, 0}
	South     = Direction{0, -1}
	West      = Direction{-1, 0}
	NorthEast = Direction{-1, 1}
	NorthWest = Direction{1, 1}
	SouthEast = Direction{-1, -1}
	SouthWest = Direction{1, -1}
)

var Directions = []Direction{
	North, East, South, West,
	NorthEast, NorthWest, SouthEast, SouthWest,
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	input := "./day04/test.txt"

	if part == 1 {
		ans := part1(input)
		fmt.Println("Answer:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Answer:", ans)
	}
}

func part1(input string) int {
	data := utils.ConvertFileToGrid(input)
	total := 0

	// Look for X's
	for x := range data {
		for y := range data[x] {
			for _, direction := range Directions {
				if getXMAS(data, [2]int{x, y}, direction) {
					// Do something if getXMAS returns true
					total++
				}
			}
		}
	}

	return total
}

func part2(input string) int {
	total := 0

	return total
}

func getXMAS(grid [][]string, xLocation [2]int, dir Direction) bool {

	return false
}
