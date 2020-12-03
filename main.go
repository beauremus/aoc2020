package main

import (
	"aoc/2020/day01"
	"aoc/2020/utilities"
	"fmt"
)

func main() {
	var day01Input, err = utilities.Read("./day01/input.txt")

	if err != nil {
		return
	}

	var numberInput []int = utilities.StringsToInts(day01Input)

	fmt.Printf("Part 1: %v\n", day01.Part1(numberInput))
	fmt.Printf("Part 2: %v\n", day01.Part2(numberInput))
}
