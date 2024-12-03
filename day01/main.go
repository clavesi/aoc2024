package main

import (
	"aoc/utils"
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// We need to create two lists
// One for the left column (a)
// One for the right column (b)
// Then sort both of them
// Then iterate over and sum the difference between a[i] and b[i]

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	input := "./day01/input.txt"

	if part == 1 {
		ans := part1(input)
		fmt.Println("Answer:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Answer:", ans)
	}
}

func part1(input string) int {
	var left []int
	var right []int

	f, err := os.Open(input)
	utils.Check(err)
	defer f.Close()

	// Go line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), "   ")
		leftVal, err := strconv.Atoi(result[0])
		rightVal, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Sort slices
	sort.Ints(left)
	sort.Ints(right)

	// Get sum of (absolute) differences
	total := 0
	for i := range len(left) {
		total += int(math.Abs(float64(left[i] - right[i])))
	}

	return total
}

func part2(input string) int {
	// Construct two hash maps
	// Both maps have the key=number and value=# of times that showed up
	left := make(map[int]int)
	right := make(map[int]int)
	f, err := os.Open(input)
	utils.Check(err)
	defer f.Close()

	// Go line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), "   ")
		leftVal, err := strconv.Atoi(result[0])
		rightVal, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		left[leftVal] = left[leftVal] + 1
		right[rightVal] = right[rightVal] + 1
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Then iterate over left map's k-v and get value from right map
	// Then it's (left_key * left_index) ^ right_val
	total := 0
	for idx, val := range left {
		total += (idx * val) * right[idx]
	}

	return total
}
