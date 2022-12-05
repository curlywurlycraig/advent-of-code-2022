package day5

import (
	"os"
	"testing"

	"crwi.uk/aoc2022/util"
)

func TestDay5(t *testing.T) {
	inputBytes, err := os.ReadFile("test.txt")
	util.NoErr(err)

	result := Day5(string(inputBytes))
	if result != "CMZ" {
		t.Errorf("Expected CMZ but was %s\n", result)
	}
}

func TestDay5Real(t *testing.T) {
	inputBytes, err := os.ReadFile("input.txt")
	util.NoErr(err)

	result := Day5(string(inputBytes))
	if result != "TDCHVHJTG" {
		t.Errorf("Expected TDCHVHJTG but was %s\n", result)
	}
}

func TestDay5Part2(t *testing.T) {
	inputBytes, err := os.ReadFile("test.txt")
	util.NoErr(err)

	result := Day5Part2(string(inputBytes))
	if result != "MCD" {
		t.Errorf("Expected MCD but was %s\n", result)
	}
}

func TestDay5Part2Real(t *testing.T) {
	inputBytes, err := os.ReadFile("input.txt")
	util.NoErr(err)

	result := Day5Part2(string(inputBytes))
	if result != "NGCMPJLHV" {
		t.Errorf("Expected NGCMPJLHV but was %s\n", result)
	}
}
