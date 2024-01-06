package main

import (
	"2022/19/part_1/objects"
	"2022/19/part_1/simulation"
	"bufio"
	"os"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func makeInitialState() simulation.State {
	return simulation.State{
		Robots:    simulation.Robots{objects.ORE_COLLECTOR: 1},
		Resources: objects.Resources{},
		Factory:   objects.Factory{IsProducing: false},
	}
}

const MINUTE_LIMIT = 24

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

	score := 0
	initialState := makeInitialState()
	for _, blueprint := range blueprints {
		blueprintScore := simulation.FindBestScore(
			blueprint.ConstructionCosts,
			initialState,
			MINUTE_LIMIT,
		)
		println("ID:", blueprint.Id, "score:", blueprintScore)
		score += blueprint.Id * blueprintScore
	}
	println("Total score:", score)
}
