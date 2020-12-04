package day02

import (
	"aoc/2020/utilities"
	"strings"
)

// Part1 takes a list of numbers and outputs a single number
func Part1(stringInput []string) int {
	validPasswordCount := 0
	passwords := utilities.Destructure(stringInput)

	for _, value := range passwords {
		letterCount := strings.Count(value.Password, value.Letter)
		if letterCount >= value.Lower && letterCount <= value.Upper {
			validPasswordCount++
		}
	}

	return validPasswordCount
}
