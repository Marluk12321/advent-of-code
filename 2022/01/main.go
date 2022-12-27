package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func shiftRightFrom(topSums *[3]int, startPos int) {
	for j := len(topSums) - 1; j > startPos; j-- {
		topSums[j] = topSums[j-1]
	}
}

func updateTopSums(topSums *[3]int, newSum int) {
	for i, sum := range topSums {
		if newSum > sum {
			shiftRightFrom(topSums, i)
			topSums[i] = newSum
			fmt.Println("Found better sum", i+1, ":", newSum)
			break
		}
	}
}

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Missing file path argument")
		return
	}

	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open", path)
		fmt.Println(err)
		return
	}
	defer f.Close()

	var sum int
	topSums := [3]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			updateTopSums(&topSums, sum)
			fmt.Println("Read sum", sum)
			sum = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Failed to read int", line)
		} else {
			sum += val
		}
	}

	fmt.Println("Top sums:", topSums)
	var total int
	for _, val := range topSums {
		total += val
	}
	fmt.Println("Total:", total)
}
