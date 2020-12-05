package day03

type strategy struct {
	Right int
	Down  int
}

// Part2 takes a list of numbers and outputs a single number
func Part2(stringInput []string) int {
	strategies := []strategy{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	var runResults []int

	for _, strat := range strategies {
		currentRow := 0
		currentColumn := 0
		treesHit := 0

		for currentRow+strat.Down < len(stringInput) {
			currentRow += strat.Down
			currentColumn += strat.Right
			effectiveColumn := currentColumn % len(stringInput[currentRow])

			if string(stringInput[currentRow][effectiveColumn]) == "#" {
				treesHit++
			}
		}

		runResults = append(runResults, treesHit)
	}

	result := 1

	for _, runResult := range runResults {
		result *= runResult
	}

	return result
}
