package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	input := "./test.txt"

	var total = part1(input)
	if total != 11 {
		t.Errorf("part1() = %v, wanted %v", total, 11)
	}
}

func Test_part2(t *testing.T) {}