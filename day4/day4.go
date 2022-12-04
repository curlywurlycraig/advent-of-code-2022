package day4

import (
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func ParseRange(v string) Range {
	values := strings.Split(v, "-")
	start, _ := strconv.Atoi(values[0])
	end, _ := strconv.Atoi(values[1])
	return Range{
		start,
		end,
	}
}

// FullyContained is true if either range is a subset of the other
func FullyContained(range1 Range, range2 Range) bool {
	firstContained := range1.start >= range2.start && range1.end <= range2.end
	secondContained := range2.start >= range1.start && range2.end <= range1.end
	return firstContained || secondContained
}

func Day4(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {
		rangeStrings := strings.Split(line, ",")
		firstRange := ParseRange(rangeStrings[0])
		secondRange := ParseRange(rangeStrings[1])
		if FullyContained(firstRange, secondRange) {
			result++
		}
	}

	return result
}

func Overlaps(range1 Range, range2 Range) bool {
	firstLess := range1.end < range2.start
	firstMore := range1.start > range2.end
	return !firstLess && !firstMore
}

func Day4Part2(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {
		rangeStrings := strings.Split(line, ",")
		firstRange := ParseRange(rangeStrings[0])
		secondRange := ParseRange(rangeStrings[1])
		if Overlaps(firstRange, secondRange) {
			result++
		}
	}

	return result
}
