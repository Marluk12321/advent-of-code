package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func idealTreeScore(grid *Grid, row int, col int) int {
	return (row * (len(grid.data) - 1 - row) *
		col * (len(grid.data[row]) - 1 - col))
}

func (grid *Grid) getTreeScore(row int, col int) int {
	if row == 0 || row == len(grid.data)-1 {
		return 0
	}
	rowHeights := grid.data[row]
	if col == 0 || col == len(rowHeights)-1 {
		return 0
	}
	colHeights := make([]int, len(grid.data))
	for i, row := range grid.data {
		colHeights[i] = row[col]
	}

	viewpointHeight := rowHeights[col]
	totalScore := 1
	for i := col - 1; i >= 0; i-- {
		if rowHeights[i] >= viewpointHeight || i == 0 {
			totalScore *= col - i
			break
		}
	}
	for i := col + 1; i < len(rowHeights); i++ {
		if rowHeights[i] >= viewpointHeight || i == len(rowHeights)-1 {
			totalScore *= i - col
			break
		}
	}
	for i := row - 1; i >= 0; i-- {
		if colHeights[i] >= viewpointHeight || i == 0 {
			totalScore *= row - i
			break
		}
	}
	for i := row + 1; i < len(colHeights); i++ {
		if colHeights[i] >= viewpointHeight || i == len(colHeights)-1 {
			totalScore *= i - row
			break
		}
	}
	return totalScore
}

func (grid *Grid) getBestTreeScore() int {
	var bestScore int
	for i, row := range grid.data {
		for j := range row {
			if idealTreeScore(grid, i, j) <= bestScore {
				continue
			}
			if score := grid.getTreeScore(i, j); score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
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

	fmt.Println("Visible trees:", grid.getBestTreeScore())
}
