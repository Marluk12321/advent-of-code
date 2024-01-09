package main

import (
	"2022/19/part_1/objects"
	"2022/19/part_2/search"
	"bufio"
	"os"
	"time"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func makeInitialState() search.State {
	return search.State{
		Robots:    search.Robots{objects.ORE_COLLECTOR: 1},
		Resources: objects.Resources{},
	}
}

const SCORING_RESOURCE = objects.OPEN_GEODE
const MINUTE_LIMIT = 32

func main() {
	blueprints := []objects.Blueprint{}

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		blueprint := objects.MakeBlueprint(line)
		blueprints = append(blueprints, blueprint)
	}
	file.Close()

	score := 1
	initialState := makeInitialState()
	startAll := time.Now()
	for _, blueprint := range blueprints[:3] {
		start := time.Now()
		blueprintScore := search.FindBestScore(
			blueprint.ConstructionCosts,
			SCORING_RESOURCE,
			MINUTE_LIMIT,
			initialState,
		)
		end := time.Now()
		println("ID:", blueprint.Id, "| score:", blueprintScore,
				"| time:", end.Sub(start).String())
		score *= blueprintScore
	}
	endAll := time.Now()
	println("Total score:", score, "| time:", endAll.Sub(startAll).String())
}
