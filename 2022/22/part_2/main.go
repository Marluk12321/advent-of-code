package main

import (
	"2022/22/part_2/instructions"
	"2022/22/part_2/mapping"
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func buildInitalState(boardMap mapping.Board) mapping.State {
	firstRow := boardMap[0]
	return mapping.State{
		Position: mapping.Position{
			Row: 0,
			Col: firstRow.Offset,
		},
		Facing: mapping.FACING_RIGHT,
	}
}

func run(cube mapping.Cube, instructions instructions.Instructions) mapping.State {
	state := buildInitalState(cube.Board)
	for _, instruction := range instructions {
		state = instruction.Apply(cube, state)
	}
	return state
}

func main() {
	board := mapping.Board{}
	var instructionList instructions.Instructions
	parsingMap := true

	cubeSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	file := openFile(os.Args[2])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		if line == "" {
			parsingMap = false
			continue
		}
		if parsingMap {
			row := mapping.BuildRow(line)
			fmt.Println("Row:", row)
			board = append(board, row)
		} else {
			instructionList = instructions.BuildInstructions(line)
			for _, instruction := range instructionList {
				fmt.Println("Instruction:", instruction)
			}
		}
	}
	file.Close()

	start := time.Now()
	cube := mapping.BuildCube(board, cubeSize)
	finalState := run(cube, instructionList)
	end := time.Now()
	password := 1000*(finalState.Position.Row+1) + 4*(finalState.Position.Col+1) + int(finalState.Facing)
	println("Password:", password, "| time:", end.Sub(start).String())
}
