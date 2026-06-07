package main

import (
	"2022/23/part_2/util"
	"2022/23/part_2/world"
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

func checkScanner(scanner *bufio.Scanner) {
	err := scanner.Err()
	if err != nil {
		panic(err)
	}
}

func parseLine(line string) []world.Position {
	parsedElves := make([]world.Position, 0)
	for col, c := range line {
		switch c {
		case '#':
			elf := world.Position{X: col, Y: 0}
			parsedElves = append(parsedElves, elf)
		case '.':
			continue
		default:
			panic(c)
		}
	}
	return parsedElves
}

func main() {
	elves := make([]world.Position, 0)

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		util.Debug(line)
		lineElves := parseLine(line)
		for i := range lineElves {
			lineElves[i].Y = row
		}
		elves = append(elves, lineElves...)
		row += 1
	}
	file.Close()
	checkScanner(scanner)
	util.Debug(elves)

	start := time.Now()
	requiredIterations := world.FindStableState(elves)
	end := time.Now()

	println("Required iterations:", requiredIterations + 1, "\nTime:", end.Sub(start).String())
}
