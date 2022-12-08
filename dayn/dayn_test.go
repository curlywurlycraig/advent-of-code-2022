package dayn

import (
	"os"
	"testing"
)

func TestDayN(t *testing.T) {
	inputBytes, _ := os.ReadFile("test.txt")
	result := DayN(string(inputBytes))

	if result != 0 {
		t.Errorf("Expected 0 but was %d\n", result)
	}
}
