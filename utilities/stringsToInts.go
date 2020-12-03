package utilities

import "strconv"

func StringsToInts(input []string) []int {
	var output []int

	for _, value := range input {
		if newValue, err := strconv.Atoi(value); err == nil {
			output = append(output, newValue)
		}
	}

	return output
}
