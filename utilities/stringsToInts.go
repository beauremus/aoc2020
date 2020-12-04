package utilities

import "strconv"

// StringsToInts takes an array of strings and returns a list of ints.
func StringsToInts(input []string) []int {
	var output []int

	for _, value := range input {
		if newValue, err := strconv.Atoi(value); err == nil {
			output = append(output, newValue)
		}
	}

	return output
}
