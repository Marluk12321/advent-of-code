package main

import (
	"2022/14/part_1/sand"
	"2022/14/part_1/world"
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

func main() {
	var rockFormations []world.RockFormation

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rockFormation := world.ParseRockFormation(line)
		rockFormations = append(rockFormations, rockFormation)
	}
	file.Close()

	sandSpawn := world.Position{X: 500, Y: 0}
	w := world.MakeWorld(rockFormations)
	sand.Fill(&w, sandSpawn)
	fmt.Println("Resting sand:", len(w.SandPositions))
}
