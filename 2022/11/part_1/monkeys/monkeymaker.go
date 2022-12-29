package monkeys

import (
	"strconv"
	"strings"
)

func toNumber(numberField string) int {
	if numberField[len(numberField)-1] == ',' {
		numberField = numberField[:len(numberField)-1]
	}
	number, err := strconv.Atoi(numberField)
	if err != nil {
		panic(err)
	}
	return number
}

func loadItems(itemDesc string) ItemQueue {
	var queue ItemQueue
	fields := strings.Fields(itemDesc)
	for _, numberField := range fields[2:] {
		queue.add(toNumber(numberField))
	}
	return queue
}

func makeOperation(operator string) func(int, int) int {
	switch operator {
	case "+":
		return func(x int, y int) int {
			return x + y
		}
	case "*":
		return func(x int, y int) int {
			return x * y
		}
	default:
		panic(operator)
	}
}

func makeUpdateFunc(updateDesc string) func(int) int {
	fields := strings.Fields(updateDesc)
	operation := makeOperation(fields[4])
	rValue := fields[5]
	if rValue == "old" {
		return func(old int) int {
			return operation(old, old)
		}
	}
	value := toNumber(rValue)
	return func(old int) int {
		return operation(old, value)
	}
}

func makeTest(testDesc string) func(int) bool {
	fields := strings.Fields(testDesc)
	modulo := toNumber(fields[3])
	return func(value int) bool {
		return value%modulo == 0
	}
}

func getValue(valueDesc string) int {
	fields := strings.Fields(valueDesc)
	return toNumber(fields[5])
}

func makeTargetFunc(targetingDesc []string) func(int) int {
	test := makeTest(targetingDesc[0])
	value1 := getValue(targetingDesc[1])
	value2 := getValue(targetingDesc[2])
	return func(value int) int {
		if test(value) {
			return value1
		} else {
			return value2
		}
	}
}

type MonkeyDesc [6]string

func MakeMonkey(description *MonkeyDesc) Monkey {
	return Monkey{
		Items:       loadItems(description[1]),
		updateWorry: makeUpdateFunc(description[2]),
		getTarget:   makeTargetFunc(description[3:6]),
	}
}
