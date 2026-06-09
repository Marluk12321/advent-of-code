package main

import (
	"2022/24/part_1/util"
	"2022/24/part_1/world"
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

func findEntryExit(line string) int {
	for col, c := range line {
		switch c {
		case '#':
			continue
		case '.':
			return col
		default:
			panic(c)
		}
	}
	panic(line)
}

func parseBlizzards(line string, row int) []world.Blizzard {
	parsedBlizzards := make([]world.Blizzard, 0)
	for col, c := range line {
		switch c {
		case '#':
			continue
		case '^':
			blizzard := world.Blizzard{Row: row, Col: col, Direction: world.UP}
			parsedBlizzards = append(parsedBlizzards, blizzard)
		case 'v':
			blizzard := world.Blizzard{Row: row, Col: col, Direction: world.DOWN}
			parsedBlizzards = append(parsedBlizzards, blizzard)
		case '<':
			blizzard := world.Blizzard{Row: row, Col: col, Direction: world.LEFT}
			parsedBlizzards = append(parsedBlizzards, blizzard)
		case '>':
			blizzard := world.Blizzard{Row: row, Col: col, Direction: world.RIGHT}
			parsedBlizzards = append(parsedBlizzards, blizzard)
		default:
			panic(c)
		}
	}
	return parsedBlizzards
}

func main() {
	valley := world.Valley{Rows: 0, Cols: 0}

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		util.Debug(line)
		if len(line) < 4 {
			panic(line)
		}

		if line[1] == '#' || line[len(line)-2] == '#' {
			entryExitCol := findEntryExit(line)
			if valley.Rows == 0 {
				valley.EntryCol = entryExitCol
			} else {
				valley.ExitCol = entryExitCol
			}
		} else {
			lineBlizzards := parseBlizzards(line, valley.Rows)
			valley.Blizzards = append(valley.Blizzards, lineBlizzards...)
		}

		valley.Rows += 1
		valley.Cols = len(line)
	}
	file.Close()
	checkScanner(scanner)
	util.Debug(valley)

	start := time.Now()
	end := time.Now()

	println("Time:", end.Sub(start).String())
}
