package day3

import (
	"strings"
)

func scoreItem(item rune) int {
	if item < 'a' {
		return int(item-'A') + 27
	}

	return int(item-'a') + 1
}

func makeCompartment(items []rune) map[rune]bool {
	result := make(map[rune]bool)
	for _, item := range items {
		result[item] = true
	}
	return result
}

func common(first []rune, second []rune) []rune {
	firstCmp := makeCompartment(first)
	secondCmp := makeCompartment(second)

	results := make([]rune, 0)
	for k := range secondCmp {
		if in := firstCmp[k]; in {
			results = append(results, k)
		}
	}
	return results
}

func Day3(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		commonItem := common([]rune(firstHalf), []rune(secondHalf))

		// Assuming always one common item
		total += scoreItem(commonItem[0])
	}

	return total
}

func Day3Part2(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for i := 0; i < len(lines); i += 3 {
		group := lines[i : i+3]

		// The item common between all three bags is the item that is
		// common between "what is common between the first two bags" and the third bag.
		badgeItem := common(
			common([]rune(group[0]), []rune(group[1])),
			[]rune(group[2]),
		)

		total += scoreItem(badgeItem[0])
	}
	return total
}
