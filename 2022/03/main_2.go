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

func findBadge(elves *[3]string) rune {
	unique2 := toSet(elves[1])
	unique3 := toSet(elves[2])
	for _, c := range elves[0] {
		if unique2[c] && unique3[c] {
			return c
		}
	}
	panic("No overlap")
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
	var elf_counter int
	var elves [3]string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		elves[elf_counter] = scanner.Text()
		if elf_counter == len(elves)-1 {
			badge := findBadge(&elves)
			fmt.Println(elves, string(badge))
			total_priority += toPriority(badge)
		}
		elf_counter = (elf_counter + 1) % len(elves)
	}

	fmt.Println("Total priority:", total_priority)
}
