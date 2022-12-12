package day12

import (
	"fmt"
	"strings"
)

type Position [2]int

type Path struct {
	stepCount int
	curr      Position
	currValue byte
}

func Day12(input string) int {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)

	paths := []Path{
		getStartPath(lines),
	}

	bestCounts := make(map[Position]int)

	completePaths := []Path{}

	for len(paths) > 0 {
		currPath, remainingPaths := paths[0], paths[1:]
		paths = remainingPaths

		if currPath.currValue == 'E' {
			completePaths = append(completePaths, currPath)
			continue
		}

		for _, neighbour := range getNeighbours(currPath, width, height) {
			neighbourValue := lines[neighbour[1]][neighbour[0]]
			realNeighbourValue := neighbourValue
			if realNeighbourValue == 'E' {
				realNeighbourValue = 'z'
			}

			if int(realNeighbourValue)-int(currPath.currValue) <= 1 {
				bestCount, ok := bestCounts[neighbour]
				if ok && bestCount <= currPath.stepCount+1 {
					continue
				}

				bestCounts[neighbour] = currPath.stepCount + 1

				paths = append(paths, Path{
					stepCount: currPath.stepCount + 1,
					curr:      neighbour,
					currValue: neighbourValue,
				})
			}
		}
	}

	winner, ok := bestCounts[getEndPosition(lines)]
	fmt.Println(winner, ok)
	minSteps := completePaths[0].stepCount
	for _, p := range completePaths {
		if p.stepCount < minSteps {
			minSteps = p.stepCount
		}
	}
	return minSteps
}

func Day12Part2(input string) int {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	height := len(lines)

	paths := getAllStartPaths(lines)

	bestCounts := make(map[Position]int)

	completePaths := []Path{}

	for len(paths) > 0 {
		currPath, remainingPaths := paths[0], paths[1:]
		paths = remainingPaths

		if currPath.currValue == 'E' {
			completePaths = append(completePaths, currPath)
			continue
		}

		for _, neighbour := range getNeighbours(currPath, width, height) {
			neighbourValue := lines[neighbour[1]][neighbour[0]]
			realNeighbourValue := neighbourValue
			if realNeighbourValue == 'E' {
				realNeighbourValue = 'z'
			}

			if int(realNeighbourValue)-int(currPath.currValue) <= 1 {
				bestCount, ok := bestCounts[neighbour]
				if ok && bestCount <= currPath.stepCount+1 {
					continue
				}

				bestCounts[neighbour] = currPath.stepCount + 1

				paths = append(paths, Path{
					stepCount: currPath.stepCount + 1,
					curr:      neighbour,
					currValue: neighbourValue,
				})
			}
		}
	}

	winner, ok := bestCounts[getEndPosition(lines)]
	fmt.Println(winner, ok)
	minSteps := completePaths[0].stepCount
	for _, p := range completePaths {
		if p.stepCount < minSteps {
			minSteps = p.stepCount
		}
	}
	return minSteps
}

func getStartPath(lines []string) Path {
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			if lines[y][x] != 'S' {
				continue
			}

			return Path{
				curr:      Position{x, y},
				currValue: 'a',
				stepCount: 0,
			}
		}
	}

	// unreachable with well formatted puzzles
	panic("couldn't find the start position")
}

func getAllStartPaths(lines []string) []Path {
	results := make([]Path, 0)

	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == 'S' || lines[y][x] == 'a' {
				results = append(results, Path{
					curr:      Position{x, y},
					currValue: 'a',
					stepCount: 0,
				})
			}
		}
	}

	return results
}

func getEndPosition(lines []string) Position {
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == 'E' {
				return Position{x, y}
			}
		}
	}

	// unreachable with well formatted puzzles
	panic("couldn't find the end position")
}

func getNeighbours(p Path, width, height int) []Position {
	x, y := p.curr[0], p.curr[1]
	results := []Position{}
	if x > 0 {
		results = append(results, Position{x - 1, y})
	}
	if y > 0 {
		results = append(results, Position{x, y - 1})
	}
	if x < width-1 {
		results = append(results, Position{x + 1, y})
	}
	if y < height-1 {
		results = append(results, Position{x, y + 1})
	}

	return results
}

func CopyMap(v map[Position]bool) map[Position]bool {
	result := make(map[Position]bool)

	for k, v := range v {
		result[k] = v
	}
	return result
}
