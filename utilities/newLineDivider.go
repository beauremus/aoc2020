package utilities

import (
	"strings"
)

// NewLiveDivider takes a string of file contents and returns groups by blank newline
func NewLiveDivider(input string) []string {
	// var output []string
	output := strings.Split(strings.ReplaceAll(input, "\n", " "), "  ")

	// for _, value := range temp {
	// 	fmt.Printf("%v\n", value)
	// }

	// fmt.Printf("%v\n", temp)

	// output = append(output, "Not dead yet!")

	return output
}
