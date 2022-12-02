package day2

import (
	"os"
	"testing"

	"crwi.uk/aoc2022/util"
)

func TestDay2(t *testing.T) {
	inputBytes, err := os.ReadFile("./test.txt")
	util.NoErr(err)

	result := Day2(string(inputBytes))
	if result != 15 {
		t.Errorf("Expected 15 but got %d\n", result)
	}
}

func TestDay2Real(t *testing.T) {
	inputBytes, err := os.ReadFile("./input.txt")
	util.NoErr(err)

	result := Day2(string(inputBytes))
	if result != 13052 {
		t.Errorf("Expected 13052 but got %d\n", result)
	}
}

func TestDay2Part2(t *testing.T) {
	inputBytes, err := os.ReadFile("./test.txt")
	util.NoErr(err)

	result := Day2Part2(string(inputBytes))
	if result != 12 {
		t.Errorf("Expected 12 but got %d\n", result)
	}
}

func TestDay2Part2Real(t *testing.T) {
	inputBytes, err := os.ReadFile("./input.txt")
	util.NoErr(err)

	result := Day2Part2(string(inputBytes))
	if result != 13693 {
		t.Errorf("Expected 13693 but got %d\n", result)
	}
}
