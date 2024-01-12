package main

import (
	"2022/20/part_1/movement"
	"bufio"
	"os"
	"strconv"
	"time"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

var OFFSETS = [3]int{1000, 2000, 3000}

func main() {
	numbers := []int{}

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	file.Close()

	start := time.Now()
	moved := movement.Move(numbers)
	sum := 0
	for _, offset := range OFFSETS {
		index := offset % len(moved)
		sum += moved[index]
	}
	end := time.Now()
	println("Sum:", sum, "| time:", end.Sub(start).String())
}
