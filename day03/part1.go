package day03

// Part1 takes a list of numbers and outputs a single number
func Part1(stringInput []string) int {
	right := 3
	down := 1
	currentRow := 0
	currentColumn := 0
	treesHit := 0

	for currentRow+down < len(stringInput) {
		currentRow += down
		currentColumn += right
		effectiveColumn := currentColumn % len(stringInput[currentRow])

		if string(stringInput[currentRow][effectiveColumn]) == "#" {
			treesHit++
		}
	}

	return treesHit
}
