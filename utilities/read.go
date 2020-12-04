package utilities

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Read takes a string of a file path and returns the fields from the file
func Read(path string) (input []string, err error) {
	data, err := ioutil.ReadFile(path)
	output := strings.Fields(string(data))

	if err != nil {
		fmt.Println("File reading error", err)
	}

	return output, err
}
