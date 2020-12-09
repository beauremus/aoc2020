package day04

import (
	"aoc/2020/utilities"
	"regexp"
	"strconv"
	"strings"
)

func between(min int, max int) func(int) bool {
	return func(testValue int) bool {
		return testValue >= min && testValue <= max
	}
}

func isHexColor(input string) bool {
	matched, _ := regexp.MatchString(`#[0-9a-f]{6}`, input)

	return matched
}

func isPID(input string) bool {
	matched, _ := regexp.MatchString(`[0-9]{9}`, input)

	return matched
}

// Part2 takes a list of numbers and outputs a single number
func Part2(stringInput []string) int {
	validPassportCount := 0
	byr := between(1920, 2002)
	iyr := between(2010, 2020)
	eyr := between(2020, 2030)
	hgtcm := between(150, 193)
	hgtin := between(59, 76)
	eclValues := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, passport := range utilities.RequiredFieldsPresent(stringInput) {
		validFieldCount := 0

		for key, value := range passport {
			if key == "ecl" {
				if !utilities.Contains(eclValues, value) {
					break
				}
			}
			if key == "pid" {
				if !isPID(value) {
					break
				}
			}
			if key == "eyr" {
				year, _ := strconv.Atoi(value)
				if !eyr(year) {
					break
				}
			}
			if key == "hcl" {
				if !isHexColor(value) {
					break
				}
			}
			if key == "byr" {
				year, _ := strconv.Atoi(value)
				if !byr(year) {
					break
				}
			}
			if key == "iyr" {
				year, _ := strconv.Atoi(value)
				if !iyr(year) {
					break
				}
			}
			if key == "hgt" {
				if strings.HasSuffix(value, "cm") {
					height, _ := strconv.Atoi(strings.TrimSuffix(value, "cm"))
					if !hgtcm(height) {
						break
					}
				} else if strings.HasSuffix(value, "in") {
					height, _ := strconv.Atoi(strings.TrimSuffix(value, "in"))
					if !hgtin(height) {
						break
					}
				}
			}

			validFieldCount++
		}

		if validFieldCount >= 7 {
			// fmt.Println(passport)
			validPassportCount++
		}
	}

	return validPassportCount
}
