package day01

import (
	"aoc/utils"
	"bufio"
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

func Test() {
	var left []int
	var right []int

	f, err := os.Open("./day01/input.txt")
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

	// Print answer
	fmt.Println(total)
}
