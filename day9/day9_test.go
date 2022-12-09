package day9

import (
	"os"
	"testing"
)

func TestDay9(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day9(string(inputBytes))

	if result != 13 {
		t.Errorf("Expected 13 but was %d\n", result)
	}
}

func TestDay9Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day9(string(inputBytes))

	if result != 6332 {
		t.Errorf("Expected 6332 but was %d\n", result)
	}
}

func TestDay9Part2(t *testing.T) {
	inputBytes, _ := os.ReadFile("test2.txt")
	result := Day9Part2(string(inputBytes))

	if result != 36 {
		t.Errorf("Expected 36 but was %d\n", result)
	}
}

func TestDay9Part2Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day9Part2(string(inputBytes))

	if result != 2511 {
		t.Errorf("Expected 2511 but was %d\n", result)
	}
}
