package world

import (
	"strconv"
	"strings"
)

func toInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return value
}

func parseRockVertex(text string) Position {
	coordinates := strings.Split(text, ",")
	return Position{
		X: toInt(coordinates[0]),
		Y: toInt(coordinates[1]),
	}
}

type RockFormation struct {
	Vertices []Position
	Min      Position
	Max      Position
}

func (formation *RockFormation) add(vertex Position) {
	formation.Vertices = append(formation.Vertices, vertex)
	if len(formation.Vertices) == 1 {
		formation.Min.X = vertex.X
		formation.Max.X = vertex.X
		formation.Min.Y = vertex.Y
		formation.Max.Y = vertex.Y
		return
	}
	if vertex.X < formation.Min.X {
		formation.Min.X = vertex.X
	} else if vertex.X > formation.Max.X {
		formation.Max.X = vertex.X
	}
	if vertex.Y < formation.Min.Y {
		formation.Min.Y = vertex.Y
	} else if vertex.Y > formation.Max.Y {
		formation.Max.Y = vertex.Y
	}
}

func ParseRockFormation(text string) RockFormation {
	coordinatePairs := strings.Split(text, " -> ")
	formation := RockFormation{
		Vertices: make([]Position, 0, len(coordinatePairs)),
	}
	for _, coordinates := range coordinatePairs {
		formation.add(parseRockVertex(coordinates))
	}
	return formation
}
