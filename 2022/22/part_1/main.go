package main

import (
	"2022/22/part_1/board"
	"bufio"
	"os"
	"strings"
	"time"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func buildInitalState(boardMap board.Map) board.State {
	firstRow := boardMap[0]
	return board.State{
		Row:    0,
		Col:    firstRow.Offset,
		Facing: board.FACING_RIGHT,
	}
}

func run(boardMap board.Map, instructions board.Instructions) board.State {
	state := buildInitalState(boardMap)
	for _, instruction := range instructions {
		state = instruction.Apply(boardMap, state)
	}
	return state
}

func main() {
	boardMap := board.Map{}
	var instructions board.Instructions
	parsingMap := true

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		if line == "" {
			parsingMap = false
			continue
		}
		if parsingMap {
			row := board.BuildRow(line)
			boardMap = append(boardMap, row)
		} else {
			instructions = board.BuildInstructions(line)
		}
	}
	file.Close()

	start := time.Now()
	finalState := run(boardMap, instructions)
	end := time.Now()
	password := 1000*(finalState.Row+1) + 4*(finalState.Col+1) + int(finalState.Facing)
	println("Password:", password, "| time:", end.Sub(start).String())
}
