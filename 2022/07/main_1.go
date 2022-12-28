package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	parent         *Directory
	files          map[string]int
	subdirectories map[string]Directory
}

func makeDirectory(parent *Directory) Directory {
	directory := Directory{parent: parent}
	directory.files = map[string]int{}
	directory.subdirectories = map[string]Directory{}
	return directory
}

func (directory *Directory) getSize() int {
	var totalSize int
	for _, size := range directory.files {
		totalSize += size
	}
	for _, subdirectory := range directory.subdirectories {
		totalSize += subdirectory.getSize()
	}
	return totalSize
}

func (directory *Directory) visit(visitor func(*Directory)) {
	visitor(directory)
	for _, subdirectory := range directory.subdirectories {
		subdirectory.visit(visitor)
	}
}

type Filesystem struct {
	root             Directory
	currentDirectory *Directory
}

func makeFilesystem() Filesystem {
	filesystem := Filesystem{root: makeDirectory(nil)}
	filesystem.currentDirectory = &filesystem.root
	return filesystem
}

func (filesystem *Filesystem) processCommand(words []string) {
	if words[0] == "cd" {
		switch words[1] {
		case "..":
			filesystem.currentDirectory = filesystem.currentDirectory.parent
		case "/":
			filesystem.currentDirectory = &filesystem.root
		default:
			subdir := filesystem.currentDirectory.subdirectories[words[1]]
			filesystem.currentDirectory = &subdir
		}
	}
}

func (filesystem *Filesystem) add(words []string) {
	currentDirectory := filesystem.currentDirectory
	if words[0] == "dir" {
		currentDirectory.subdirectories[words[1]] = makeDirectory(currentDirectory)
	} else {
		size, err := strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}
		currentDirectory.files[words[1]] = size
	}
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

	filesystem := makeFilesystem()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		words := strings.Fields(line)
		if words[0] == "$" {
			filesystem.processCommand(words[1:])
		} else {
			filesystem.add(words[:])
		}
	}

	var sum int
	filesystem.root.visit(func(dir *Directory) {
		if size := dir.getSize(); size <= 100000 {
			sum += size
		}
	})
	fmt.Println("Sum:", sum)
}
