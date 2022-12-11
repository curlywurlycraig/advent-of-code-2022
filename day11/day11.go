package day11

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	mult int = iota
	add

	part1Rounds = 20
	part2Rounds = 10000
)

type Monkey struct {
	items               []int
	operationType       int
	operationValue      int
	operationValueIsOld bool
	testDivide          int
	onTrueThrow         int
	onFalseThrow        int

	inspectCount int
}

func parseMonkey(input string) Monkey {
	result := Monkey{}

	lines := strings.Split(input, "\n")

	result.items = parseItems(lines[1])

	operationType, operationValue, valueIsOld := parseOperation(lines[2])
	result.operationType = operationType
	result.operationValue = operationValue
	result.operationValueIsOld = valueIsOld

	result.testDivide = parseTest(lines[3])
	result.onTrueThrow = parseTrueThrow(lines[4])
	result.onFalseThrow = parseFalseThrow(lines[5])

	result.inspectCount = 0

	return result
}

// Example:
// `  Starting items: 85, 77, 77`
func parseItems(itemsLine string) []int {
	itemsSep := strings.TrimPrefix(itemsLine, "  Starting items: ")
	itemStrings := strings.Split(itemsSep, ", ")
	items := []int{}
	for _, itemString := range itemStrings {
		item, _ := strconv.Atoi(itemString)
		items = append(items, item)
	}
	return items
}

// Example:
// `  Operation: new = old * 7`
func parseOperation(operationLine string) (int, int, bool) {
	operationString := strings.TrimPrefix(operationLine, "  Operation: new = old ")
	operationParts := strings.Split(operationString, " ")
	var operation int
	if operationParts[0] == "*" {
		operation = mult
	} else {
		operation = add
	}

	operationValueString, _ := strconv.Atoi(operationParts[1])

	isOld := false
	if strings.HasSuffix(operationLine, "old") {
		isOld = true
	}

	return operation, operationValueString, isOld
}

// Example:
// `  Test: divisible by 19`
func parseTest(testLine string) int {
	testValueString := strings.TrimPrefix(testLine, "  Test: divisible by ")
	testValue, _ := strconv.Atoi(testValueString)
	return testValue
}

// Example:
// `    If true: throw to monkey 6`
func parseTrueThrow(trueLine string) int {
	valueString := strings.TrimPrefix(trueLine, "    If true: throw to monkey ")
	value, _ := strconv.Atoi(valueString)
	return value
}

// Example:
// `    If false: throw to monkey 6`
func parseFalseThrow(falseLine string) int {
	valueString := strings.TrimPrefix(falseLine, "    If false: throw to monkey ")
	value, _ := strconv.Atoi(valueString)
	return value
}

func Day11(input string, rounds int, amCalm bool) int {
	monkeyStrings := strings.Split(input, "\n\n")
	monkeys := make([]*Monkey, 0)

	// greatest modulo is simply the number that, when applying as modulo,
	// still results in correct divisability of numbers for all monkeys.
	// I don't know what this is called in maths.
	// It is the product of all divisors
	greatestModulo := 1
	for _, monkeyString := range monkeyStrings {
		newMonkey := parseMonkey(monkeyString)
		monkeys = append(monkeys, &newMonkey)
		greatestModulo = greatestModulo * newMonkey.testDivide
	}

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				// apply worry operation
				newWorry := item
				operand := monkey.operationValue
				if monkey.operationValueIsOld {
					operand = item
				}

				if monkey.operationType == mult {
					newWorry = item * operand
				} else {
					newWorry = item + operand
				}

				if amCalm {
					newWorry = int(math.Floor(float64(newWorry) / 3.0))
				} else {
					newWorry = newWorry % greatestModulo
				}

				// check division test
				throwTo := monkey.onFalseThrow
				if newWorry%monkey.testDivide == 0 {
					throwTo = monkey.onTrueThrow
				}

				// throw to one of the monkeys
				monkeys[throwTo].items = append(monkeys[throwTo].items, newWorry)

				monkey.inspectCount++
			}

			// reset monkeys items (they have all been thrown)
			monkey.items = []int{}
		}
	}

	inspectCounts := []int{}
	for _, monkey := range monkeys {
		inspectCounts = append(inspectCounts, monkey.inspectCount)
	}
	sort.Ints(inspectCounts)
	fmt.Println(inspectCounts)

	return inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2]
}
