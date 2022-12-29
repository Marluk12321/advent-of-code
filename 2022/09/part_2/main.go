package main

import (
	rope "2022/09/part_2/rope"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PositionSet map[rope.Position]bool

func (positionSet *PositionSet) add(position *rope.Position) {
	(*positionSet)[*position] = true
}

func parse(line *string) (rope.Direction, int) {
	parts := strings.Fields(*line)
	var direction rope.Direction
	switch parts[0] {
	case "U":
		direction = rope.UP
	case "D":
		direction = rope.DOWN
	case "L":
		direction = rope.LEFT
	case "R":
		direction = rope.RIGHT
	default:
		panic(parts[0])
	}
	repeats, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return direction, repeats
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

	rope := rope.MakeRope(10)
	tailPositions := PositionSet{}
	tailPositions.add(&rope.Tail().Position)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction, repeats := parse(&line)
		for i := 0; i < repeats; i++ {
			rope.Move(direction)
			tailPositions.add(&rope.Tail().Position)
		}
	}

	fmt.Println("Tail positions:", len(tailPositions))
}
