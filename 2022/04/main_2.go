package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toRange(elf string) [2]int {
	var result [2]int
	rangeStr := strings.Split(elf, "-")
	for i, s := range rangeStr {
		var err error
		result[i], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return result
}

func overlap(range1 *[2]int, range2 *[2]int) bool {
	return range1[0] <= range2[1] && range1[1] >= range2[0]
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

	var overlap_count int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		elves := strings.Split(line, ",")
		range1 := toRange(elves[0])
		range2 := toRange(elves[1])
		if overlap(&range1, &range2) {
			fmt.Println(line, true)
			overlap_count++
		} else {
			fmt.Println(line, false)
		}
	}

	fmt.Println("overlap:", overlap_count)
}
