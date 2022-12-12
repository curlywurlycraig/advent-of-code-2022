package day12

import (
	"os"
	"testing"
)

func TestDay12(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day12(string(inputBytes))

	if result != 31 {
		t.Errorf("Expected 31 but was %d\n", result)
	}
}

func TestDay12Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day12(string(inputBytes))

	if result != 394 {
		t.Errorf("Expected 394 but was %d\n", result)
	}
}

func TestDay12Part2(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day12Part2(string(inputBytes))

	if result != 29 {
		t.Errorf("Expected 29 but was %d\n", result)
	}
}

func TestDay12Part2Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day12Part2(string(inputBytes))

	if result != 388 {
		t.Errorf("Expected 388 but was %d\n", result)
	}
}
