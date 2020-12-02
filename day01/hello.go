package main

import (
    "fmt"
    "strings"
    "io/ioutil"
    "strconv"
)

func read() (input []string, err error) {
    data, err := ioutil.ReadFile("input.txt")
    output := strings.Fields(string(data))

    if err != nil {
        fmt.Println("File reading error", err)
    }

    return output, err
}

func stringsToInts(input []string) []int {
    var output []int

    for _, value := range input {
        if newValue, err := strconv.Atoi(value); err == nil {
            output = append(output, newValue)
        }
    }

    return output
}

func main() {
    var input, err = read()

    if err != nil { return }



    var numberInput []int = stringsToInts(input)

    for _, firstValue := range numberInput {
        for _, secondValue := range numberInput {
            if firstValue + secondValue == 2020 {
                fmt.Printf("%v\n", firstValue * secondValue)
                return
            }
        }
    }
}
