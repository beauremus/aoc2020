package utilities

import (
	"strconv"
	"strings"
)

// PasswordTest represents the structure of AOC input
type PasswordTest struct {
	Lower    int
	Upper    int
	Letter   string
	Password string
}

// Destructure takes a string array and outputs a slice of slices
func Destructure(stringInput []string) []PasswordTest {
	var output []PasswordTest
	test := 0

	for test < len(stringInput) {
		first := strings.Split((stringInput[test]), "-")
		lower, _ := strconv.Atoi(first[0])
		upper, _ := strconv.Atoi(first[1])
		second := stringInput[test+1]
		letter := string(second[0])
		third := stringInput[test+2]
		password := string(third)

		output = append(output, PasswordTest{lower, upper, letter, password})

		test += 3
	}

	return output
}
