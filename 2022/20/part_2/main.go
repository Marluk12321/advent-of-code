package main

import (
	"2022/20/part_2/mixing"
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

const DECRYPTION_KEY = 811589153
const MIXING_REPEATS = 10
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
		numbers = append(numbers, number * DECRYPTION_KEY)
	}
	file.Close()

	start := time.Now()
	cycleLinks := mixing.ToCycleLinks(numbers)
	for i := 0; i < MIXING_REPEATS; i++ {
		mixing.Mix(cycleLinks)
	}

	zeroStartingSlice := mixing.ToZeroStartingSlice(cycleLinks)
	sum := 0
	for _, offset := range OFFSETS {
		index := offset % len(zeroStartingSlice)
		sum += zeroStartingSlice[index]
	}
	end := time.Now()
	println("Sum:", sum, "| time:", end.Sub(start).String())
}
