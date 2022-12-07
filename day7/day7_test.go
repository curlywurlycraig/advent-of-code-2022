package day7

import (
	"os"
	"testing"
)

func TestDay7(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day7(string(inputBytes))

	if result != 95437 {
		t.Errorf("Expected 95437 but was %d\n", result)
	}
}

func TestDay7Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day7(string(inputBytes))

	if result != 1390824 {
		t.Errorf("Expected 1390824 but was %d\n", result)
	}
}

func TestDay7Part2(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := Day7Part2(string(inputBytes))

	if result != 24933642 {
		t.Errorf("Expected 24933642 but was %d\n", result)
	}
}

func TestDay7Part2Real(t *testing.T) {
	inputBytes, _ := os.ReadFile("input.txt")
	result := Day7Part2(string(inputBytes))

	if result != 7490863 {
		t.Errorf("Expected 7490863 but was %d\n", result)
	}
}
