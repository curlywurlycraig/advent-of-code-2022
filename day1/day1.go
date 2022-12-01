package day1

import (
	"sort"
	"strconv"
	"strings"

	"crwi.uk/aoc2022/util"
)

// Day1 returns an int representing the max calories
func Day1(input string) int {
	elfSections := strings.Split(input, "\n\n")
	max := 0

	for _, section := range elfSections {
		calorieLines := strings.Split(section, "\n")
		elfTotal := 0
		for _, calorieLine := range calorieLines {
			calorieAsInt, err := strconv.Atoi(calorieLine)
			util.NoErr(err)
			elfTotal += calorieAsInt
		}

		if elfTotal > max {
			max = elfTotal
		}
	}

	return max
}

func Day1Part2(input string) int {
	elfSections := strings.Split(input, "\n\n")
	totals := make([]int, len(elfSections))

	for elfIndex, section := range elfSections {
		calorieLines := strings.Split(section, "\n")
		elfTotal := 0
		for _, calorieLine := range calorieLines {
			calorieAsInt, err := strconv.Atoi(calorieLine)
			util.NoErr(err)
			elfTotal += calorieAsInt
		}

		totals[elfIndex] = elfTotal
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	return totals[0] + totals[1] + totals[2]
}
