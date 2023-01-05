package objects

import (
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (position Position) DistanceTo(other Position) int {
	return abs(position.X-other.X) + abs(position.Y-other.Y)
}

func toInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return value
}

func MakePosition(text string) Position {
	positionParts := strings.Split(text, ", ")
	return Position{
		X: toInt(positionParts[0][2:]),
		Y: toInt(positionParts[1][2:]),
	}
}
