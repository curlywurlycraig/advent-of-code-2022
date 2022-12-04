package day4

import (
	"os"
	"testing"

	"crwi.uk/aoc2022/util"
)

func TestDay4(t *testing.T) {
	inputBytes, err := os.ReadFile("test.txt")
	util.NoErr(err)

	result := Day4(string(inputBytes))
	if result != 2 {
		t.Errorf("expected 2 but was %d\n", result)
	}
}

func TestDay4Real(t *testing.T) {
	inputBytes, err := os.ReadFile("input.txt")
	util.NoErr(err)

	result := Day4(string(inputBytes))
	if result != 424 {
		t.Errorf("expected 424 but was %d\n", result)
	}
}

func TestDay4Part2(t *testing.T) {
	inputBytes, err := os.ReadFile("test.txt")
	util.NoErr(err)

	result := Day4Part2(string(inputBytes))
	if result != 4 {
		t.Errorf("expected 4 but was %d\n", result)
	}
}

func TestDay4Part2Real(t *testing.T) {
	inputBytes, err := os.ReadFile("input.txt")
	util.NoErr(err)

	result := Day4Part2(string(inputBytes))
	if result != 804 {
		t.Errorf("expected 804 but was %d\n", result)
	}
}
