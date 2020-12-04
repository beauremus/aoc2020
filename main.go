package main

import (
	"aoc/2020/day01"
	"aoc/2020/day02"
	"aoc/2020/utilities"
	"fmt"
)

func main() {
	var day01Input, day01Err = utilities.Read("./day01/input.txt")
	var day02Input, day02Err = utilities.Read("./day02/input.txt")

	if day01Err != nil || day02Err != nil {
		return
	}

	fmt.Printf(
		"Day 1 - Part 1: %v\n",
		day01.Part1(utilities.StringsToInts(day01Input)),
	)
	fmt.Printf(
		"Day 1 - Part 2: %v\n",
		day01.Part2(utilities.StringsToInts(day01Input)),
	)
	fmt.Printf("Day 2 - Part 1: %v\n", day02.Part1(day02Input))
	fmt.Printf("Day 2 - Part 2: %v\n", day02.Part2(day02Input))
}
