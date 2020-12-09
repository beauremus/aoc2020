package main

import (
	"aoc/2020/day01"
	"aoc/2020/day02"
	"aoc/2020/day03"
	"aoc/2020/day04"
	"aoc/2020/utilities"
	"fmt"
)

func main() {
	var day01Input, day01Err = utilities.Read(
		utilities.Params{Path: "./day01/input.txt"})
	var day02Input, day02Err = utilities.Read(
		utilities.Params{Path: "./day02/input.txt"})
	var day03Input, day03Err = utilities.Read(
		utilities.Params{Path: "./day03/input.txt"})
	var day04Input, day04Err = utilities.Read(
		utilities.Params{Path: "./day04/input.txt", Mod: utilities.NewLiveDivider})

	if day01Err != nil ||
		day02Err != nil ||
		day03Err != nil ||
		day04Err != nil {
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
	fmt.Printf("Day 3 - Part 1: %v\n", day03.Part1(day03Input))
	fmt.Printf("Day 3 - Part 2: %v\n", day03.Part2(day03Input))
	fmt.Printf("Day 4 - Part 1: %v\n", day04.Part1(day04Input))
	fmt.Printf("Day 4 - Part 2: %v\n", day04.Part2(day04Input))
}
