package utilities

import (
	"strings"
)

// NewLiveDivider takes a string of file contents and returns groups by blank newline
func NewLiveDivider(input string) []string {
	output := strings.Split(strings.ReplaceAll(input, "\n", " "), "  ")

	return output
}
