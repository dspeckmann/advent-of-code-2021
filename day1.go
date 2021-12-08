package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func getMeasurements(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		panic("Error while opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	measurements := make([]int, 0)
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Invalid value found in input file.")
		}
		measurements = append(measurements, measurement)
	}

	file.Close()

	return measurements
}

func partOne() int {
	measurements := getMeasurements("day1-part1.txt")

	increases := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increases++
		}
	}

	return increases
}

func getSum(measurements []int, index int) int {
	return measurements[index] + measurements[index-1] + measurements[index-2]
}

func partTwo() int {
	measurements := getMeasurements("day1-part2.txt")

	increases := 0
	for i := 3; i < len(measurements); i++ {
		if getSum(measurements, i) > getSum(measurements, i-1) {
			increases++
		}
	}

	return increases
}
