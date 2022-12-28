package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	row int
	col int
}

type PositionGenerator struct {
	initialPosition Position
	rowStep         int
	colStep         int
}

func (generator *PositionGenerator) generate(offset int) Position {
	return Position{
		row: generator.initialPosition.row + offset*generator.rowStep,
		col: generator.initialPosition.col + offset*generator.colStep,
	}
}

type Grid struct {
	data [][]int
}

func (grid *Grid) addRow(line *string) {
	row := make([]int, 0, len(*line))
	for _, c := range *line {
		height, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		row = append(row, height)
	}
	grid.data = append(grid.data, row)
}

func getVisiblePositions(sequence *[]int, positionGenerator *PositionGenerator) []Position {
	visiblePositions := []Position{}
	tallestVisible := -1
	for i, height := range *sequence {
		if height > tallestVisible {
			tallestVisible = height
			position := positionGenerator.generate(i)
			visiblePositions = append(visiblePositions, position)
		}
	}
	return visiblePositions
}

func reverse(sequence *[]int) []int {
	reversed := make([]int, 0, len(*sequence))
	for i := len(*sequence) - 1; i >= 0; i-- {
		reversed = append(reversed, (*sequence)[i])
	}
	return reversed
}

func (grid *Grid) getColumns() [][]int {
	columns := make([][]int, len(grid.data))
	for _, row := range grid.data {
		for i, value := range row {
			columns[i] = append(columns[i], value)
		}
	}
	return columns
}

func update(position_set *map[Position]bool, positions []Position) {
	for _, position := range positions {
		(*position_set)[position] = true
	}
}

func (grid *Grid) countVisibleTrees() int {
	visiblePositions := map[Position]bool{}
	for i, row := range grid.data {
		rightwardGenerator := PositionGenerator{
			initialPosition: Position{i, 0},
			colStep:         1,
		}
		rightwardPositions := getVisiblePositions(&row, &rightwardGenerator)
		update(&visiblePositions, rightwardPositions)

		reveresedRow := reverse(&row)
		leftwardGenerator := PositionGenerator{
			initialPosition: Position{i, len(row) - 1},
			colStep:         -1,
		}
		leftwardPositions := getVisiblePositions(&reveresedRow, &leftwardGenerator)
		update(&visiblePositions, leftwardPositions)
	}

	for i, column := range grid.getColumns() {
		downwardGenerator := PositionGenerator{
			initialPosition: Position{0, i},
			rowStep:         1,
		}
		downwardPositions := getVisiblePositions(&column, &downwardGenerator)
		update(&visiblePositions, downwardPositions)

		reveresedCol := reverse(&column)
		upwardGenerator := PositionGenerator{
			initialPosition: Position{len(grid.data) - 1, i},
			rowStep:         -1,
		}
		upwardPositions := getVisiblePositions(&reveresedCol, &upwardGenerator)
		update(&visiblePositions, upwardPositions)
	}

	return len(visiblePositions)
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

	grid := Grid{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		grid.addRow(&line)
	}

	fmt.Println("Visible trees:", grid.countVisibleTrees())
}
