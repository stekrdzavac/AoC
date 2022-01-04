package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)


func readInput() (string, error) {
	fileContents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return "", err
	}
	trimmedContents := strings.Trim(string(fileContents), "\n")

	return trimmedContents, nil
}

func main() {

	input, err := readInput()
	if err != nil {
		fmt.Printf("got error: %+v", err)
	}

}
