package day10

import (
	"os"
	"testing"
)

// Doesn't actually complete enough cycles to get a valid answer,
// but useful for debugging.
func TestDay10Tiny(t *testing.T) {
	inputBytes, _ := os.ReadFile("tinytest.txt")
	result := Day10(string(inputBytes))

	if result != 0 {
		t.Errorf("Expected 0 but was %d\n", result)
	}
}

func TestDay10(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day10(string(inputBytes))

	if result != 13140 {
		t.Errorf("Expected 13140 but was %d\n", result)
	}
}

func TestDay10Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day10(string(inputBytes))

	if result != 13480 {
		t.Errorf("Expected 13480 but was %d\n", result)
	}
}

func TestDay10Part2(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	expectedOutputBytes, _ := os.ReadFile("part2testoutput.txt")

	result := Day10Part2(string(inputBytes))

	if result != string(expectedOutputBytes) {
		t.Errorf("Output does not match\n")
	}
}

func TestDay10Part2Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	expectedOutputBytes, _ := os.ReadFile("part2realoutput.txt")

	result := Day10Part2(string(inputBytes))

	if result != string(expectedOutputBytes) {
		t.Errorf("Output does not match\n")
	}
}
