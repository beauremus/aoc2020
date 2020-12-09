package day04

import (
	"aoc/2020/utilities"
	"regexp"
	"strconv"
	"strings"
)

var (
	hexRE = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	pidRE = regexp.MustCompile(`^[0-9]{9}$`)
)

func between(min int, max int) func(int) bool {
	return func(testValue int) bool {
		return testValue >= min && testValue <= max
	}
}

func isHexColor(input string) bool {
	return hexRE.MatchString(input)
}

func isPID(input string) bool {
	return pidRE.MatchString(input)
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

	passport:
		for key, value := range passport {
			switch key {
			case "ecl":
				if !utilities.Contains(eclValues, value) {
					break passport
				}
			case "pid":
				if !isPID(value) {
					break passport
				}
			case "eyr":
				year, _ := strconv.Atoi(value)
				if !eyr(year) {
					break passport
				}
			case "hcl":
				if !isHexColor(value) {
					break passport
				}
			case "byr":
				year, _ := strconv.Atoi(value)
				if !byr(year) {
					break passport
				}
			case "iyr":
				year, _ := strconv.Atoi(value)
				if !iyr(year) {
					break passport
				}
			case "hgt":
				if strings.HasSuffix(value, "cm") {
					height, _ := strconv.Atoi(strings.TrimSuffix(value, "cm"))
					if !hgtcm(height) {
						break passport
					}
				} else if strings.HasSuffix(value, "in") {
					height, _ := strconv.Atoi(strings.TrimSuffix(value, "in"))
					if !hgtin(height) {
						break passport
					}
				}
			case "cid":
				break passport
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
