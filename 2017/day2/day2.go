package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput() [] string {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	var txtLines []string
    for scanner.Scan() {
		txtLines = append(txtLines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return  txtLines
}


func main() {

	input := readInput()

	fmt.Println("Part one =", partOne(input))
	fmt.Println("Part one =", partTwo(input))

}


func partOne(input [] string) int {

	checkSum := 0

	for _, row := range input {

		max, min := 0, math.MaxInt64

		words := strings.Split(row, "\t")

		for _, word := range words {

			num, err := strconv.Atoi(string(word))
			if err != nil {
				log.Fatal(err)
			}

			if num < min {
				min = num
			}

			if num > max {
				max = num
			}
		}

		checkSum += max - min
	}

	return checkSum
}


func partTwo(input [] string) int {

	checkSum := 0

	for _, row := range input {

		words := strings.Split(row, "\t")

		for i := 0; i < len(words);i++{

			num, err := strconv.Atoi(string(words[i]))
			if err != nil {
				log.Fatal(err)
			}

			for j := i + 1 ; j < len(words); j++ {
				nextNum, err := strconv.Atoi(string(words[j]))
				if err != nil {
					log.Fatal(err)
				}

				if num >= nextNum && num % nextNum == 0 {
					checkSum += num / nextNum
					break
				}
				if nextNum > num && nextNum % num == 0  {
					checkSum += nextNum / num
					break
				}
			}
		}
	}

	return checkSum
}
