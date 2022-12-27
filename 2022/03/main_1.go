package main

import (
	"bufio"
	"fmt"
	"os"
)

func toSet(compartment string) map[rune]bool {
	set := map[rune]bool{}
	for _, c := range compartment {
		set[c] = true
	}
	return set
}

func findOverlap(compartment1 string, compartment2 string) string {
	unique_c1 := toSet(compartment1)
	unique_c2 := toSet(compartment2)
	overlapped := []rune{}
	for c := range unique_c1 {
		if unique_c2[c] {
			overlapped = append(overlapped, c)
		}
	}
	return string(overlapped)
}

func toPriority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	return int(c-'A') + 27
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

	var total_priority int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		compartment1, compartment2 := line[:len(line)/2], line[len(line)/2:]
		overlap := findOverlap(compartment1, compartment2)
		fmt.Println(compartment1, compartment2, overlap)
		for _, c := range overlap {
			total_priority += toPriority(c)
		}
	}

	fmt.Println("Total priority:", total_priority)
}
