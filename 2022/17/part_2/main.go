package main

import (
	"2022/17/part_2/objects"
	"2022/17/part_2/simulation"
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
	var text string

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()
	}
	file.Close()

	roomSize := 7
	spawnLimit := 1000000000000

	room := objects.MakeRoom(roomSize)
	blockTypes := &objects.BlockTypes
	jetPattern := simulation.MakeJetPattern(text)
	simulation.Simulate(&room, blockTypes, &jetPattern, spawnLimit)
	fmt.Println("Total height:", room.GetHeight())
}
