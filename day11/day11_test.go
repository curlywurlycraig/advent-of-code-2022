package day11

import (
	"os"
	"testing"
)

func TestDay11(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day11(string(inputBytes), part1Rounds, true)

	if result != 10605 {
		t.Errorf("Expected 10605 but was %d\n", result)
	}
}

func TestDay11Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day11(string(inputBytes), part1Rounds, true)

	if result != 54752 {
		t.Errorf("Expected 54752 but was %d\n", result)
	}
}

func TestDay11Part2(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day11(string(inputBytes), part2Rounds, false)

	if result != 2713310158 {
		t.Errorf("Expected 2713310158 but was %d\n", result)
	}
}

func TestDay11Part2Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day11(string(inputBytes), part2Rounds, false)

	if result != 13606755504 {
		t.Errorf("Expected 13606755504 but was %d\n", result)
	}
}
