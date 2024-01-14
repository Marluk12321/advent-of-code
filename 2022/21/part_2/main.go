package main

import (
	"2022/21/part_2/calculation"
	"2022/21/part_2/tree"
	"bufio"
	"os"
	"time"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

const ROOT_NAME = "root"
const HUMAN_NAME = "humn"

func main() {
	nodeDescs := map[string]tree.NodeDesc{}

	file := openFile(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		desc := tree.BuildNodeDesc(line)
		nodeDescs[desc.NodeName] = desc
	}
	file.Close()

	start := time.Now()
	humanValue := calculation.CalculateVariableValue(nodeDescs, ROOT_NAME, HUMAN_NAME)
	end := time.Now()
	println("Human value:", humanValue, "| time:", end.Sub(start).String())
}
