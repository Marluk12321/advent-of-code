package main

import (
	"2022/17/part_1/objects"
	"2022/17/part_1/simulation"
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
	var jetPattern string

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jetPattern = scanner.Text()
	}
	file.Close()

	roomSize := 7
	spawnLimit := 2022

	blockTypes := &objects.BlockTypes
	room := objects.MakeRoom(roomSize)
	simulation.Simulate(&room, blockTypes, jetPattern, spawnLimit)
	fmt.Println("Total height:", room.GetHeight())
}
