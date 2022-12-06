package day6

func Day6(input string) int {
	for i := 4; i < len(input); i++ {
		if !anySame(input[i-4 : i]) {
			return i
		}
	}

	panic("(unreachable) couldn't find any string with all different characters")
}

func Day6Part2(input string) int {
	// I was right lol
	for i := 14; i < len(input); i++ {
		if !anySame(input[i-14 : i]) {
			return i
		}
	}

	panic("(unreachable) couldn't find any string with all different characters")
}

func anySame(input string) bool {
	// Map is overkill for part 1, but I am suspecting that part 2 will
	// increase from 4 characters
	foundMap := make(map[rune]bool)

	for _, c := range input {
		_, found := foundMap[c]
		if found {
			return true
		}
		foundMap[c] = true
	}

	return false
}
