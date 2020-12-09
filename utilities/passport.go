package utilities

import "strings"

// Contains takes a source slice and compares each value to the test returning a bool
func Contains(source []string, test string) bool {
	for _, value := range source {
		if value == test {
			return true
		}
	}

	return false
}

// RequiredFieldsPresent takes a list of passports and output a list of valid passports
func RequiredFieldsPresent(passports []string) []map[string]string {
	requiredFields := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}
	var output []map[string]string

	for _, passport := range passports {
		passportFields := strings.Split(passport, " ")
		passportMap := make(map[string]string)

		if len(passportFields) >= 7 {
			validFieldCount := 0

			for _, field := range passportFields {
				parsedField := strings.Split(field, ":")
				passportMap[parsedField[0]] = parsedField[1]

				if Contains(requiredFields, parsedField[0]) {
					validFieldCount++
				}
			}

			if validFieldCount >= 7 {
				output = append(output, passportMap)
			}
		}
	}

	return output
}
