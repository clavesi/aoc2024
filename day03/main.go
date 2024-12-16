package main

import (
	"aoc/utils"
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	input := "./day03/input.txt"

	if part == 1 {
		ans := part1(input)
		fmt.Println("Answer:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Answer:", ans)
	}
}

func part1(input string) int {
	total := 0

	// Open file
	f, err := os.Open(input)
	utils.Check(err)
	defer f.Close()

	// REGEX
	r, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")

	// Go line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Get only proper substrings
		mulList := r.FindAllString(scanner.Text(), -1)
		for _, mul := range mulList {
			// Trim to get just the numbers
			s := strings.Replace(mul, "mul(", "", -1)
			s = strings.Replace(s, ")", "", -1)
			nums := strings.Split(s, ",")

			// Convert and add to total
			first, _ := strconv.Atoi(nums[0])
			second, _ := strconv.Atoi(nums[1])
			total += first * second
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return total
}

func part2(input string) int {
	fmt.Println(input)
	return 0
}
