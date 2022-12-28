package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ParserState int

const (
	READING_STACKS ParserState = iota
	AWAITING_COMMANDS
	READING_COMMANDS
)

func push(stack *[]rune, c rune) {
	*stack = append((*stack), c)
}

func popSlice(stack *[]rune, size int) []rune {
	index := len(*stack) - size
	slice := (*stack)[index:]
	*stack = (*stack)[:index]
	return slice
}

func readStackLine(line *string, reverseStacks *[][]rune) {
	for i, c := range *line {
		if i%4 == 1 && c != ' ' {
			stackIndex := i / 4
			for len(*reverseStacks) < stackIndex+1 {
				*reverseStacks = append(*reverseStacks, []rune{})
			}
			push(&(*reverseStacks)[stackIndex], c)
		}
	}
}

func fillStacks(reverseStacks *[][]rune, stacks *[][]rune) {
	for _, reverseStack := range *reverseStacks {
		stack := make([]rune, 0, len(reverseStack))
		for i := len(reverseStack) - 1; i >= 0; i-- {
			push(&stack, reverseStack[i])
		}
		(*stacks) = append((*stacks), stack)
	}
}

type Command struct {
	quantity  int
	fromIndex int
	toIndex   int
}

func parseCommand(line *string) Command {
	parts := strings.Fields(*line)
	quantity, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	from, err := strconv.Atoi(parts[3])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(parts[5])
	if err != nil {
		panic(err)
	}
	return Command{quantity, from - 1, to - 1}
}

func performCommand(command *Command, stacks *[][]rune) {
	fromStack := &(*stacks)[command.fromIndex]
	toStack := &(*stacks)[command.toIndex]
	slice := popSlice(fromStack, command.quantity)
	*toStack = append(*toStack, slice...)
}

func readCommandLine(line *string, stacks *[][]rune) {
	command := parseCommand(line)
	performCommand(&command, stacks)
}

func printSizes(stacks *[][]rune) {
	sizes := make([]int, 0, len(*stacks))
	for _, stack := range *stacks {
		sizes = append(sizes, len(stack))
	}
	fmt.Println("Sizes:", sizes)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing file path argument")
		fmt.Println(os.Args)
		return
	}

	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open", path)
		fmt.Println(err)
		return
	}
	defer f.Close()

	var reverseStacks [][]rune
	var stacks [][]rune

	parserState := READING_STACKS
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		switch parserState {
		case READING_STACKS:
			if line[0] == '[' {
				readStackLine(&line, &reverseStacks)
			} else {
				fillStacks(&reverseStacks, &stacks)
				parserState = AWAITING_COMMANDS

				fmt.Println("Stacks:")
				for _, stack := range stacks {
					fmt.Println(stack)
				}
				fmt.Println()
				printSizes(&stacks)
			}
		case AWAITING_COMMANDS:
			if len(line) > 0 && line[0] == 'm' {
				readCommandLine(&line, &stacks)
				parserState = READING_COMMANDS
				fmt.Println(line)
				printSizes(&stacks)
			}
		case READING_COMMANDS:
			readCommandLine(&line, &stacks)
			fmt.Println(line)
			printSizes(&stacks)
		}
	}

	var stackTops []rune
	for _, stack := range stacks {
		stackTops = append(stackTops, stack[len(stack)-1])
	}
	fmt.Println("Stack tops:", string(stackTops))
}
