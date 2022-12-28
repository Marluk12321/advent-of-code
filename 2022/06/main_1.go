package main

import (
	"bufio"
	"fmt"
	"os"
)

func push(queue *[]byte, char byte) {
	(*queue) = append((*queue), char)
}

func pop(queue *[]byte) byte {
	char := (*queue)[0]
	(*queue) = (*queue)[1:]
	return char
}

func allUnique(queue *[]byte) bool {
	set := map[byte]bool{}
	for _, char := range *queue {
		set[char] = true
	}
	return len(set) == 4
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing file path argument")
		fmt.Println(os.Args)
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

	var counter int
	queue := make([]byte, 0, 4)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		char := scanner.Text()
		fmt.Println(char)
		push(&queue, char[0])
		counter++
		if len(queue) == 4 {
			if allUnique(&queue) {
				break
			}
			pop(&queue)
		}
	}

	fmt.Println("Chars read:", counter)
}
