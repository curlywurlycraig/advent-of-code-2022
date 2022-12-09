package day9

import (
	"strconv"
	"strings"
)

type Position [2]int

func Day9(input string) int {
	lines := strings.Split(input, "\n")

	headPosition := Position{0, 0}
	tailPosition := Position{0, 0}
	visitedPositions := make(map[Position]bool)
	visitedPositions[tailPosition] = true

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		for i := 0; i < amount; i++ {
			moveHead(&headPosition, direction)
			moveTail(&headPosition, &tailPosition)
			visitedPositions[tailPosition] = true
		}
	}

	return len(visitedPositions)
}

func moveHead(head *Position, direction string) {
	switch direction {
	case "R":
		head[0]++
	case "L":
		head[0]--
	case "U":
		head[1]--
	case "D":
		head[1]++
	}
}

func moveTail(head *Position, tail *Position) {
	if !touching(head, tail) && head[0] != tail[0] && head[1] != tail[1] {
		// diagonal case
		if head[0] > tail[0] {
			tail[0]++
		} else {
			tail[0]--
		}

		if head[1] > tail[1] {
			tail[1]++
		} else {
			tail[1]--
		}
	} else {
		if head[0]-tail[0] > 1 {
			tail[0]++
		}
		if head[0]-tail[0] < -1 {
			tail[0]--
		}
		if head[1]-tail[1] > 1 {
			tail[1]++
		}
		if head[1]-tail[1] < -1 {
			tail[1]--
		}
	}
}

func touching(head *Position, tail *Position) bool {
	return intAbs(head[0]-tail[0]) < 2 && intAbs(head[1]-tail[1]) < 2
}

func intAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func Day9Part2(input string) int {
	lines := strings.Split(input, "\n")

	positions := [10]Position{}
	visitedPositions := make(map[Position]bool)
	visitedPositions[positions[9]] = true

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		for i := 0; i < amount; i++ {
			moveHead(&positions[0], direction)
			for posIdx, _ := range positions {
				if posIdx == 0 {
					// skip head
					continue
				}

				moveTail(&positions[posIdx-1], &positions[posIdx])
			}
			visitedPositions[positions[9]] = true
		}
	}

	return len(visitedPositions)
}
