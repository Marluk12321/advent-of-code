package main

import (
	"2022/15/part_2/objects"
	"2022/15/part_2/search"
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
		sensor.SetBeacon(objects.MakeBeacon(parts[1]))
		sensors = append(sensors, sensor)
	}
	file.Close()

	start := 0
	end := 4000000
	for y := start; y <= end; y++ {
		knownSegments := search.FindKnownSegments(sensors, y)
		if len(knownSegments) > 1 {
			beaconPosition := objects.Position{
				X: knownSegments[0].End + 1,
				Y: y,
			}
			fmt.Println("Beacon location:", beaconPosition)
			frequency := 4000000*beaconPosition.X + beaconPosition.Y
			fmt.Println("Frequency:", frequency)
			break
		}
	}
}
