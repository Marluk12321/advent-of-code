package main

import (
	system "2022/10/part_2/system"
	"bufio"
	"fmt"
	"os"
)

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

	system := system.MakeSystem()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		system.CPU.LoadInstruction(line)
	}
	f.Close()

	for len(system.CPU.PendingOperations) > 0 {
		system.StartCycle()
		system.EndCycle()
	}
	system.CRT.Print()
}
