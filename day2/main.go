package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Movement struct {
	Direction string
	Distance  int
}

func main() {
	var movements []Movement
	for _, row := range strings.Split(input, "\n") {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}

		vals := strings.Split(row, " ")
		direction := vals[0]
		distance, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}
		movements = append(movements, Movement{Direction: direction, Distance: distance})
	}

	totalDistance := part1(movements)
	fmt.Printf("Distance 1: %d\n", totalDistance)

	totalDistance = part2(movements)
	fmt.Printf("Distance 2: %d\n", totalDistance)
}

func part1(movements []Movement) (totalDistance int) {
	var (
		horizontalPosition int
		depth              int
	)
	for _, movement := range movements {
		switch movement.Direction {
		case "forward":
			horizontalPosition += movement.Distance
		case "down":
			depth += movement.Distance
		case "up":
			depth -= movement.Distance
		}
	}
	return horizontalPosition * depth
}

func part2(movements []Movement) (totalDistance int) {
	var (
		aim                int
		depth              int
		horizontalPosition int
	)
	for _, movement := range movements {
		switch movement.Direction {
		case "forward":
			horizontalPosition += movement.Distance
			depth += aim * movement.Distance
		case "down":
			aim += movement.Distance
		case "up":
			aim -= movement.Distance
		}
	}

	return horizontalPosition * depth
}
