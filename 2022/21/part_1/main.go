package main

import (
	"2022/21/part_1/tree"
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
	root := tree.BuildTree(nodeDescs, ROOT_NAME)
	rootValue := root.GetValue()
	end := time.Now()
	println("Root value:", rootValue, "| time:", end.Sub(start).String())
}
