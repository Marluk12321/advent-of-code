package main

import (
	"2022/16/part_2/optimization"
	"2022/16/part_2/valve"
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
	var valves []valve.Valve

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		valve := valve.MakeValve(line)
		valves = append(valves, valve)
	}
	file.Close()

	startValve := "AA"
	availableTime := 26
	pressureReleased := optimization.MaxPressureReleased(valves, startValve, availableTime)
	fmt.Println("Pressure released:", pressureReleased)
}
