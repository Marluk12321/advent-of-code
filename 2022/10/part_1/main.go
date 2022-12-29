package main

import (
	cpu "2022/10/part_1/cpu"
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

	cpu := cpu.MakeCPU()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cpu.LoadInstruction(line)
	}
	f.Close()

	var sum int
	for len(cpu.PendingOperations) > 0 {
		cpu.StartCycle()
		if cpu.Cycle >= 20 && (cpu.Cycle-20)%40 == 0 {
			sum += cpu.Cycle * cpu.X
		}
		cpu.EndCycle()
	}
	fmt.Println("Last cycle:", cpu.Cycle)
	fmt.Println("Total:", sum)
}
