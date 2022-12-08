package day8

import (
	"strings"
)

type position [2]int

func Day8(input string) int {
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])

	visible := make(map[position]bool)

	// left to right
	for v := 0; v < height; v++ {
		max := -1
		for h := 0; h < width; h++ {
			height := heightAt(lines, v, h)
			if height > max {
				max = height
				visible[position{v, h}] = true
			}
		}
	}

	// Right to left
	for v := 0; v < height; v++ {
		max := -1
		for h := width - 1; h >= 0; h-- {
			height := heightAt(lines, v, h)
			if height > max {
				max = height
				visible[position{v, h}] = true
			}
		}
	}

	// Top to bottom
	for h := 0; h < width; h++ {
		max := -1
		for v := 0; v < height; v++ {
			height := heightAt(lines, v, h)
			if height > max {
				max = height
				visible[position{v, h}] = true
			}
		}
	}

	// Bottom to top
	for h := 0; h < width; h++ {
		max := -1
		for v := height - 1; v >= 0; v-- {
			height := heightAt(lines, v, h)
			if height > max {
				max = height
				visible[position{v, h}] = true
			}
		}
	}

	return len(visible)
}

func heightAt(input []string, v int, h int) int {
	return int(input[v][h] - '0')
}

func Day8Part2(input string) int {
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])

	max := -1

	for v := 0; v < height; v++ {
		for h := 0; h < width; h++ {
			left := leftScore(lines, v, h)
			right := rightScore(lines, v, h)
			up := upScore(lines, v, h)
			down := downScore(lines, v, h)
			totalScore := left * right * up * down
			if totalScore > max {
				max = totalScore
			}
		}
	}

	return max
}

func leftScore(lines []string, v int, h int) int {
	treeHeight := heightAt(lines, v, h)
	result := 0
	for h2 := h - 1; h2 >= 0; h2-- {
		result++
		if treeHeight <= heightAt(lines, v, h2) {
			return result
		}
	}

	return result
}

func rightScore(lines []string, v int, h int) int {
	width := len(lines[0])
	treeHeight := heightAt(lines, v, h)
	result := 0
	for h2 := h + 1; h2 < width; h2++ {
		result++
		if treeHeight <= heightAt(lines, v, h2) {
			return result
		}
	}

	return result
}

func downScore(lines []string, v int, h int) int {
	height := len(lines)
	treeHeight := heightAt(lines, v, h)
	result := 0
	for v2 := v + 1; v2 < height; v2++ {
		result++
		if treeHeight <= heightAt(lines, v2, h) {
			return result
		}
	}

	return result
}

func upScore(lines []string, v int, h int) int {
	treeHeight := heightAt(lines, v, h)
	result := 0
	for v2 := v - 1; v2 >= 0; v2-- {
		result++
		if treeHeight <= heightAt(lines, v2, h) {
			return result
		}
	}

	return result
}
