package day01

// Part2 takes a list of numbers and outputs a single number
func Part2(numberInput []int) int {
	for _, firstValue := range numberInput {
		for _, secondValue := range numberInput {
			for _, thirdValue := range numberInput {
				if firstValue+secondValue+thirdValue == 2020 {
					return firstValue * secondValue * thirdValue
				}
			}
		}
	}

	return -1
}
