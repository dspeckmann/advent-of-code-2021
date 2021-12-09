package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(partOne())
}

func getLines() []string {
	file, err := os.Open("day3.txt")
	if err != nil {
		panic("Error while opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines
}

func partOne() int64 {
	lines := getLines()
	ones := make([]int, 0)
	zeros := make([]int, 0)
	for _, line := range lines {
		for pos, char := range line {
			if len(ones) < pos+1 {
				ones = append(ones, 0)
				zeros = append(zeros, 0)
			}
			if char == '1' {
				ones[pos]++
			} else if char == '0' {
				zeros[pos]++
			} else {
				panic("Invalid value found in line.")
			}
		}
	}

	var gammaString strings.Builder
	var epsilonString strings.Builder
	for i := 0; i < len(ones); i++ {
		if ones[i] > zeros[i] {
			gammaString.WriteByte('1')
			epsilonString.WriteByte('0')
		} else {
			gammaString.WriteByte('0')
			epsilonString.WriteByte('1')
		}
	}

	gamma, err := strconv.ParseInt(gammaString.String(), 2, 64)
	if err != nil {
		panic("Invalid gamma string.")
	}

	epsilon, err := strconv.ParseInt(epsilonString.String(), 2, 64)
	if err != nil {
		panic("Invalid epsilon string.")
	}

	return gamma * epsilon
}
