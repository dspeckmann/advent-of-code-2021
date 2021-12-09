package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Distance  int
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func getCommands() []Command {
	file, err := os.Open("day2.txt")
	if err != nil {
		panic("Error while opening file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	commands := make([]Command, 0)
	for scanner.Scan() {
		rawCommand := strings.Split(scanner.Text(), " ")
		if len(rawCommand) < 2 {
			panic("Invalid line found in input file.")
		}
		direction := rawCommand[0]
		distance, err := strconv.Atoi(rawCommand[1])
		if err != nil {
			panic("Invalid distance found in input file.")
		}
		commands = append(commands, Command{Direction: direction, Distance: distance})
	}

	file.Close()

	return commands
}

func partOne() int {
	depth := 0
	pos := 0
	commands := getCommands()
	for _, command := range commands {
		switch command.Direction {
		case "up":
			depth -= command.Distance
		case "down":
			depth += command.Distance
		case "forward":
			pos += command.Distance
		}
	}

	return depth * pos
}

func partTwo() int {
	aim := 0
	depth := 0
	pos := 0
	commands := getCommands()
	for _, command := range commands {
		switch command.Direction {
		case "up":
			aim -= command.Distance
		case "down":
			aim += command.Distance
		case "forward":
			pos += command.Distance
			depth += aim * command.Distance
		}
	}

	return depth * pos
}
