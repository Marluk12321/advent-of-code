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

func buildInitalState() mapping.State {
	return mapping.State{
		Face:         mapping.FRONT_FACE,
		FacePosition: mapping.Position{Row: 0, Col: 0},
		Facing:       mapping.FACING_RIGHT,
	}
}

func run(cube mapping.Cube, instructions instructions.Instructions) mapping.State {
	state := buildInitalState()
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
	faceOnBoard := cube.Faces[finalState.Face]
	finalPosition := mapping.Position{
		Row: faceOnBoard.Position.Row + finalState.FacePosition.Row,
		Col: faceOnBoard.Position.Col + finalState.FacePosition.Col,
	}
	password := 1000*(finalPosition.Row+1) + 4*(finalPosition.Col+1) + int(finalState.Facing)
	println("Password:", password, "| time:", end.Sub(start).String())
}
