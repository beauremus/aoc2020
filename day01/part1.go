package day01

// Part1 takes a list of numbers and outputs a single number
func Part1(numberInput []int) int {
	for _, firstValue := range numberInput {
		for _, secondValue := range numberInput {
			if firstValue+secondValue == 2020 {
				return firstValue * secondValue
			}
		}
	}

	return -1
}
