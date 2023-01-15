package main

import (
	"2022/18/part_1/calculation"
	"2022/18/part_1/objects"
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
	var cubes []objects.Cube

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cube := objects.MakeCube(line)
		cubes = append(cubes, cube)
	}
	file.Close()

	totalSurface := calculation.CalcTotalSurface(cubes)
	fmt.Println("Total surface:", totalSurface)
}
