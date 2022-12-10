package day10

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day10(input string) int {
	cycle := 0
	X := 1
	strengths := [500]int{}
	xValues := [500]int{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			strengths[cycle] = X * (cycle + 1)
			xValues[cycle] = X
			cycle++
		} else if strings.HasPrefix(line, "addx") {
			strengths[cycle] = X * (cycle + 1)
			xValues[cycle] = X
			cycle++
			strengths[cycle] = X * (cycle + 1)
			xValues[cycle] = X
			cycle++

			amount, _ := strconv.Atoi(line[5:])
			X += amount
		}
	}

	strengths[cycle] = X * (cycle + 1)
	xValues[cycle] = X

	return strengths[20-1] + strengths[60-1] + strengths[100-1] + strengths[140-1] + strengths[180-1] + strengths[220-1]
}

func Day10Part2(input string) string {
	X := 1
	xValues := []int{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			xValues = append(xValues, X)
		} else if strings.HasPrefix(line, "addx") {
			xValues = append(xValues, X)
			xValues = append(xValues, X)

			amount, _ := strconv.Atoi(line[5:])
			X += amount
		}
	}

	result := ""
	for i, xValue := range xValues {
		currX := i % 40
		if math.Abs(float64(xValue-currX)) < 2 {
			result += "#"
		} else {
			result += " "
		}

		if currX == 39 {
			result += "\n"
		}
	}

	fmt.Println(result)

	return result
}
