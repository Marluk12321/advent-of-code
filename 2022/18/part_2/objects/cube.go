package objects

import (
	"strconv"
	"strings"
)

type Cube Position

func toInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return value
}

func MakeCube(desc string) Cube {
	parts := strings.Split(desc, ",")
	return Cube{
		X: toInt(parts[0]),
		Y: toInt(parts[1]),
		Z: toInt(parts[2]),
	}
}
