package day02

import (
	"strconv"
	"strings"
)

// Part1 takes a list of numbers and outputs a single number
func Part1(stringInput []string) int {
	test := 0
	validPasswordCount := 0

	for test < len(stringInput) {
		first := strings.Split((stringInput[test]), "-")
		lower, _ := strconv.Atoi(first[0])
		upper, _ := strconv.Atoi(first[1])
		second := stringInput[test+1]
		letter := string(second[0])
		third := stringInput[test+2]
		password := string(third)
		letterCount := strings.Count(password, letter)

		if letterCount >= lower && letterCount <= upper {
			validPasswordCount++
		}

		test += 3
	}

	return validPasswordCount
}
