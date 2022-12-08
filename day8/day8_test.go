package day8

import (
	"os"
	"testing"
)

func TestDay8(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day8(string(inputBytes))

	if result != 21 {
		t.Errorf("Expected 21 but was %d\n", result)
	}
}

func TestDay8Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day8(string(inputBytes))

	if result != 1816 {
		t.Errorf("Expected 1816 but was %d\n", result)
	}
}

func TestDay8Part2(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day8Part2(string(inputBytes))

	if result != 8 {
		t.Errorf("Expected 8 but was %d\n", result)
	}
}

func TestDay8Part2Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day8Part2(string(inputBytes))

	if result != 383520 {
		t.Errorf("Expected 383520 but was %d\n", result)
	}
}
