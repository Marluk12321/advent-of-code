package main

import (
	monkeylib "2022/11/part_1/monkeys"
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

func calcScore(inspections *[]int) int {
	var best, secondBest int
	if (*inspections)[0] > (*inspections)[1] {
		best = (*inspections)[0]
		secondBest = (*inspections)[1]
	} else {
		best = (*inspections)[0]
		secondBest = (*inspections)[1]
	}
	for _, count := range (*inspections)[2:] {
		if count > best {
			secondBest = best
			best = count
		} else if count > secondBest {
			secondBest = count
		}
	}
	return best * secondBest
}

func main() {
	monkeys := monkeylib.Monkeys{}

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	var monkeyDesc monkeylib.MonkeyDesc
	var counter int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		monkeyDesc[counter] = line
		counter++
		if counter == len(monkeyDesc) {
			monkey := monkeylib.MakeMonkey(&monkeyDesc)
			monkeys.Add(&monkey)
			counter = 0
		}
	}
	file.Close()

	inspections := make([]int, len(monkeys.List))
	for monkeys.PlayedRounds < 20 {
		currentMonkey := &monkeys.List[monkeys.Current]
		inspections[monkeys.Current] += len(currentMonkey.Items)
		monkeys.TakeTurn()
	}
	fmt.Println("Score:", calcScore(&inspections))
}
