package day02

import (
	"aoc/2020/utilities"
)

// Part2 takes a list of numbers and outputs a single number
func Part2(stringInput []string) int {
	validPasswordCount := 0
	passwords := utilities.Destructure(stringInput)

	for _, value := range passwords {
		letter := value.Letter[0]
		firstLetter := value.Password[value.Lower-1]
		secondLetter := value.Password[value.Upper-1]

		if (letter == firstLetter || letter == secondLetter) &&
			firstLetter != secondLetter {
			validPasswordCount++
		}
	}

	return validPasswordCount
}
