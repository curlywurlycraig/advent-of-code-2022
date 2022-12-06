package day6

import (
	"os"
	"testing"
)

func TestDay6(t *testing.T) {
	tests := []struct {
		input     string
		expected  int
		expected2 int
	}{
		{
			input:     "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected:  7,
			expected2: 19,
		},
		{
			input:     "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected:  5,
			expected2: 23,
		},
		{
			input:     "nppdvjthqldpwncqszvftbrmjlhg",
			expected:  6,
			expected2: 23,
		},
		{
			input:     "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected:  10,
			expected2: 29,
		},
		{
			input:     "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected:  11,
			expected2: 26,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := Day6(test.input)
			if result != test.expected {
				t.Errorf("expected %d but was %d\n", test.expected, result)
			}

			result2 := Day6Part2(test.input)
			if result2 != test.expected2 {
				t.Errorf("expected %d but was %d\n", test.expected, result2)
			}
		})
	}
}

func TestDay6Real(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	result := Day6(string(input))

	if result != 1080 {
		t.Errorf("expected 1080 but was %d\n", result)
	}
}

func TestDay6Part2Real(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	result := Day6Part2(string(input))

	if result != 3645 {
		t.Errorf("expected 3645 but was %d\n", result)
	}
}
