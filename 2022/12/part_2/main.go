package main

import (
	"2022/12/part_2/search"
	"2022/12/part_2/terrain"
	"bufio"
	"fmt"
	"os"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func copyHeighMap(terrain *terrain.Terrain) [][]rune {
	result := make([][]rune, len(terrain.HeightMap))
	for i, row := range terrain.HeightMap {
		rowCopy := make([]rune, len(row))
		copy(rowCopy, row)
		result[i] = rowCopy
	}
	return result
}

func getStarts(t *terrain.Terrain) []terrain.Position {
	var starts []terrain.Position
	for i, row := range t.HeightMap {
		for j, height := range row {
			if height == 'a' {
				starts = append(starts, terrain.Position{
					Row: i,
					Col: j,
				})
			}
		}
	}
	return starts
}

func toChar(current terrain.Position, next terrain.Position) rune {
	if current.Row < next.Row {
		return 'v'
	}
	if current.Row > next.Row {
		return '^'
	}
	if current.Col < next.Col {
		return '>'
	}
	return '<'
}

func printMap(runeMap *[][]rune) {
	for _, row := range *runeMap {
		fmt.Println(string(row))
	}
}

func draw(terrain *terrain.Terrain, path *[]terrain.Position) {
	mapCopy := copyHeighMap(terrain)
	start := (*path)[0]
	mapCopy[start.Row][start.Col] = 'S'
	for i, position := range (*path)[1 : len(*path)-1] {
		char := toChar(position, (*path)[i+2])
		mapCopy[position.Row][position.Col] = char
	}
	end := (*path)[len(*path)-1]
	mapCopy[end.Row][end.Col] = 'E'
	printMap(&mapCopy)
}

func main() {
	var t terrain.Terrain

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		t.AddRow(line)
	}
	file.Close()

	var overallShortestPath []terrain.Position
	shortestStartDistance := t.GetRows() * t.GetCols()
	for _, start := range getStarts(&t) {
		t.Start = start
		shortestPath, err := search.AStar(&t)
		if err != nil {
			continue
		}
		pathLength := len(shortestPath) - 1
		if pathLength < shortestStartDistance {
			shortestStartDistance = pathLength
			overallShortestPath = shortestPath
		}
	}
	draw(&t, &overallShortestPath)
	fmt.Println("Least steps:", shortestStartDistance)
}
