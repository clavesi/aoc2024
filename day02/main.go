package main

import (
	"aoc/utils"
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	input := "./day02/input.txt"

	if part == 1 {
		ans := part1(input)
		fmt.Println("Answer:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Answer:", ans)
	}
}

func part1(input string) int {
	totalSafe := 0
	// Open file
	f, err := os.Open(input)
	utils.Check(err)
	defer f.Close()

	// Go line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		levelsStr := strings.Split(scanner.Text(), " ")
		levels := make([]int, len(levelsStr))

		// Convert levels to int
		for i, s := range levelsStr {
			levels[i], err = strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
		}

		// Check condition 1
		// - Row is in increasing or decreasing order
		isAscending := sort.SliceIsSorted(levels, func(a, b int) bool {
			return levels[a] < levels[b]
		})
		isDescending := sort.SliceIsSorted(levels, func(a, b int) bool {
			return levels[a] > levels[b]
		})

		// Check condition 2
		// - Any adjacent levels differ by at least one and at most 3
		valid := true
		if isAscending || isDescending {
			for idx := 0; idx < len(levels)-1; idx++ {
				dif := utils.Abs(levels[idx] - levels[idx+1])
				if !(idx < len(levels)-1 && dif > 0 && dif < 4) {
					valid = false
					break
				}
			}
			if valid {
				totalSafe++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return totalSafe
}

func part2(input string) int {

	return -1
}
