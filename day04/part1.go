package day04

import "aoc/2020/utilities"

// Part1 takes a list of numbers and outputs a single number
func Part1(stringInput []string) int {
	return len(utilities.RequiredFieldsPresent(stringInput))
}
