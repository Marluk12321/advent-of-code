package main

import (
	"2022/15/part_1/objects"
	"2022/15/part_1/search"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func main() {
	var sensors []objects.Sensor

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		sensor := objects.MakeSensor(parts[0])
		sensor.ClosestBeacon = objects.MakeBeacon(parts[1])
		sensors = append(sensors, sensor)
	}
	file.Close()

	lineY := 2000000
	emptySpaces := search.KnownEmptySpaces(sensors, lineY)
	fmt.Println("Empty spaces:", emptySpaces)
}
