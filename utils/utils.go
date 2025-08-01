package utils

import (
	"io"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(x int) int {
	if x <= 0 {
		return -x
	}
	return x
}

func OpenFile(file string) *os.File {
	f, err := os.Open(file)
	Check(err)
	return f
}

func ReadFile(file string) []byte {
	f, err := os.Open(file)
	Check(err)
	defer f.Close()

	content, err := io.ReadAll(f)
	Check(err)
	return content
}

func ConvertFileToGrid(input string) [][]string {
	var output [][]string

	data := ReadFile(input)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		// TODO: I might need to add a delimiter argument to the this function for this split, instead of "",
		// in case another day's rows has characters or stuff separated by like spaces or commas.
		// Maybe even for the first delimiter?
		// I'll make wahtever change when needed
		output = append(output, strings.Split(line, ""))
	}

	return output
}
