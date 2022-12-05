package day5

import (
	"strconv"
	"strings"
)

func Day5(input string) string {
	parts := strings.Split(input, "\n\n")
	initialState := parts[0]
	stacks := parseInitialState(initialState)

	instructions := strings.Split(parts[1], "\n")
	for _, instruction := range instructions {
		amount, src, dst := parseInstruction(instruction)
		for i := 0; i < amount; i++ {
			srcStack := stacks[src]
			valueToMove := srcStack[len(srcStack)-1]
			stacks[dst] = append(stacks[dst], valueToMove)
			stacks[src] = srcStack[:len(srcStack)-1]
		}
	}

	result := ""
	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}
	return result
}

func Day5Part2(input string) string {
	parts := strings.Split(input, "\n\n")
	initialState := parts[0]
	stacks := parseInitialState(initialState)

	instructions := strings.Split(parts[1], "\n")
	for _, instruction := range instructions {
		amount, src, dst := parseInstruction(instruction)
		srcStack := stacks[src]
		stacks[dst] = append(stacks[dst], srcStack[len(srcStack)-amount:]...)
		stacks[src] = srcStack[:len(srcStack)-amount]
	}

	result := ""
	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}
	return result
}

func parseInstruction(instruction string) (int, int, int) {
	withoutMove := instruction[5:]
	fromSplit := strings.Split(withoutMove, " from ")
	amount, _ := strconv.Atoi(fromSplit[0])
	toSplit := strings.Split(fromSplit[1], " to ")
	src, _ := strconv.Atoi(toSplit[0])
	dst, _ := strconv.Atoi(toSplit[1])

	return amount, src - 1, dst - 1 // Parse as 0 indexed instead of 1 indexed
}

func parseInitialState(initialState string) [][]rune {
	// each stack is three characters wide plus a space, except the last one.
	numberStacks := getStackCount(initialState)

	results := [][]rune{}
	for i := 0; i < numberStacks; i++ {
		newStack := parseStack(initialState, i)
		results = append(results, newStack)
	}

	return results
}

func getStackCount(initialState string) int {
	lines := strings.Split(initialState, "\n")
	columnNumberLines := lines[len(lines)-1]

	// each stack is three characters wide plus a space, except the last one.
	return (len(columnNumberLines) + 1) / 4
}

func parseStack(initialState string, stackIdx int) []rune {
	result := make([]rune, 0)
	lines := strings.Split(initialState, "\n")

	// All but the numbers
	for _, line := range lines[:len(lines)-1] {
		runeIdx := 1 + 4*stackIdx
		currRune := line[runeIdx]
		if currRune == ' ' {
			continue
		}

		// Prepend
		result = append(
			[]rune{rune(currRune)},
			result...,
		)
	}

	return result
}
