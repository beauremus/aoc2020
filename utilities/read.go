package utilities

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Params is a path and modification function
type Params struct {
	Path string
	Mod  func(string) []string
}

// Read takes a string of a file path and returns the fields from the file
func Read(p Params) (input []string, err error) {
	data, err := ioutil.ReadFile(p.Path)
	var output []string

	if p.Mod != nil {
		output = p.Mod(string(data))
	} else {
		output = strings.Fields(string(data))
	}

	if err != nil {
		fmt.Println("File reading error", err)
	}

	return output, err
}
