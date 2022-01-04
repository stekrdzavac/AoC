package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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
	fmt.Println("Part one =", partOne(input))
	fmt.Println("Part two =", partTwo(input))

}

func partOne(input string) int {

	input = input + string(input[0])
	sum := 0
	i := 0
	for index := 0; index < len(input); {
		for i = index + 1; i < len(input); i++ {
			if (input[index] == input[i]) {
				num, _ := strconv.Atoi(string(input[index]))
				sum += num
			} else {
				break
			}
		}
		index = i
	}

	return sum
}

func partTwo(input string) int {

	sum := 0
	step := len(input) / 2

	for index := 0; index < len(input) / 2; index++ {
		secondIndex := index + step
		if input[index] == input[secondIndex] {
			num, _ := strconv.Atoi(string(input[index]))
			sum += num * 2
		}
	}

	return sum
}
