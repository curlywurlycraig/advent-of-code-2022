package day3

import (
	"os"
	"testing"

	"crwi.uk/aoc2022/util"
)

func TestScoreItem(t *testing.T) {
	tests := []struct {
		input    rune
		expected int
	}{
		{
			input:    'a',
			expected: 1,
		},
		{
			input:    'z',
			expected: 26,
		},
		{
			input:    'A',
			expected: 27,
		},
		{
			input:    'Z',
			expected: 52,
		},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			res := scoreItem(test.input)
			if res != test.expected {
				t.Errorf("Expected %d but was %d\n", test.expected, res)
			}
		})
	}
}

func TestDay3(t *testing.T) {
	inputBytes, err := os.ReadFile("test.txt")
	util.NoErr(err)

	result := Day3(string(inputBytes))
	if result != 157 {
		t.Errorf("Expected 157 but was %d\n", result)
	}
}

func TestDay3Real(t *testing.T) {
	inputBytes, err := os.ReadFile("input.txt")
	util.NoErr(err)

	result := Day3(string(inputBytes))
	if result != 7716 {
		t.Errorf("Expected 7716 but was %d\n", result)
	}
}

func TestDay3Part2(t *testing.T) {
	inputBytes, err := os.ReadFile("test.txt")
	util.NoErr(err)

	result := Day3Part2(string(inputBytes))
	if result != 70 {
		t.Errorf("Expected 70 but was %d\n", result)
	}
}

func TestDay3Part2Real(t *testing.T) {
	inputBytes, err := os.ReadFile("input.txt")
	util.NoErr(err)

	result := Day3Part2(string(inputBytes))
	if result != 2973 {
		t.Errorf("Expected 2973 but was %d\n", result)
	}
}
