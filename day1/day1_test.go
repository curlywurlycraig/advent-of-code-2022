package day1

import (
	"os"
	"testing"

	"crwi.uk/aoc2022/util"
)

func TestDay1(t *testing.T) {
	testBytes, err := os.ReadFile("test.txt")
	util.NoErr(err)

	testString := string(testBytes)
	result := Day1(testString)
	if result != 24000 {
		t.Errorf("Expected 24000 but was %d\n", result)
	}
}

func TestDay1Real(t *testing.T) {
	testBytes, err := os.ReadFile("input.txt")
	util.NoErr(err)

	testString := string(testBytes)
	result := Day1(testString)
	if result != 74198 {
		t.Errorf("Expected 74198 but was %d\n", result)
	}
}

func TestDay1Part2(t *testing.T) {
	testBytes, err := os.ReadFile("test.txt")
	util.NoErr(err)

	testString := string(testBytes)
	result := Day1Part2(testString)
	if result != 45000 {
		t.Errorf("Expected 45000 but was %d\n", result)
	}
}

func TestDay1Part2Real(t *testing.T) {
	testBytes, err := os.ReadFile("input.txt")
	util.NoErr(err)

	testString := string(testBytes)
	result := Day1Part2(testString)
	if result != 209914 {
		t.Errorf("Expected 209914 but was %d\n", result)
	}
}
