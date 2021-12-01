package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var measurements []int
	for _, row := range strings.Split(input, "\n") {
		measurement, err := strconv.Atoi(strings.TrimSpace(row))
		if err != nil {
			continue
		}
		measurements = append(measurements, measurement)
	}

	increases1 := part1(measurements)
	fmt.Println(increases1)

	increases2 := part2(measurements)
	fmt.Println(increases2)
}

func part1(measurements []int) int {
	var increases int
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increases++
		}
	}
	return increases
}

func part2(measurements []int) int {
	var increases2 int
	for i := 1; i < len(measurements)-2; i++ {
		before1 := i - 1
		middle1 := i
		after1 := i + 1

		before2 := before1 + 1
		middle2 := middle1 + 1
		after2 := after1 + 1

		window1 := measurements[before1] + measurements[middle1] + measurements[after1]
		window2 := measurements[before2] + measurements[middle2] + measurements[after2]
		if window2 > window1 {
			increases2++
		}
	}
	return increases2
}
