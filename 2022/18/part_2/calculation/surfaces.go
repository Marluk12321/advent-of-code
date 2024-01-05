package calculation

import (
	"2022/18/part_2/objects"
)

func minPos(pos1, pos2 objects.Position) objects.Position {
	resultPos := objects.Position{}
	if pos1.X < pos2.X {
		resultPos.X = pos1.X
	} else {
		resultPos.X = pos2.X
	}
	if pos1.Y < pos2.Y {
		resultPos.Y = pos1.Y
	} else {
		resultPos.Y = pos2.Y
	}
	if pos1.Z < pos2.Z {
		resultPos.Z = pos1.Z
	} else {
		resultPos.Z = pos2.Z
	}
	return resultPos
}

func maxPos(pos1, pos2 objects.Position) objects.Position {
	resultPos := objects.Position{}
	if pos1.X > pos2.X {
		resultPos.X = pos1.X
	} else {
		resultPos.X = pos2.X
	}
	if pos1.Y > pos2.Y {
		resultPos.Y = pos1.Y
	} else {
		resultPos.Y = pos2.Y
	}
	if pos1.Z > pos2.Z {
		resultPos.Z = pos1.Z
	} else {
		resultPos.Z = pos2.Z
	}
	return resultPos
}

func getBounds(cubes []objects.Cube) (objects.Position, objects.Position) {
	min := objects.Position(cubes[0])
	max := min
	for _, cube := range cubes[1:] {
		cubePosition := objects.Position(cube)
		min = minPos(min, cubePosition)
		max = maxPos(max, cubePosition)
	}
	return objects.Position{X: min.X - 1, Y: min.Y - 1, Z: min.Z - 1},
		   objects.Position{X: max.X + 1, Y: max.Y + 1, Z: max.Z + 1}
}

func getNeighbors(pos objects.Position) [6]objects.Position {
	return [6]objects.Position{
		{X: pos.X - 1, Y: pos.Y, Z: pos.Z},
		{X: pos.X + 1, Y: pos.Y, Z: pos.Z},
		{X: pos.X, Y: pos.Y - 1, Z: pos.Z},
		{X: pos.X, Y: pos.Y + 1, Z: pos.Z},
		{X: pos.X, Y: pos.Y, Z: pos.Z - 1},
		{X: pos.X, Y: pos.Y, Z: pos.Z + 1},
	}
}

func withinBounds(pos, min, max objects.Position) bool {
	return pos.X >= min.X && pos.X <= max.X &&
		   pos.Y >= min.Y && pos.Y <= max.Y &&
		   pos.Z >= min.Z && pos.Z <= max.Z
}

func countReachableSurfaces(
	cubePositions map[objects.Position]bool,
	min, max objects.Position,
) int {
	surfaceCount := 0
	toProcess := []objects.Position{min}
	seen := map[objects.Position]bool{min: true}
	for len(toProcess) > 0 {
		current := toProcess[0]
		toProcess = toProcess[1:]
		for _, neighbor := range getNeighbors(current) {
			if cubePositions[neighbor] {
				surfaceCount++
			} else if !seen[neighbor] && withinBounds(neighbor, min, max) {
				toProcess = append(toProcess, neighbor)
				seen[neighbor] = true
			}
		}
	}
	return surfaceCount
}

func CalcTotalSurface(cubes []objects.Cube) int {
	if len(cubes) == 0 {
		return 0
	}

	cubePositions := map[objects.Position]bool{}
	for _, cube := range cubes {
		cubePosition := objects.Position(cube)
		cubePositions[cubePosition] = true
	}

	min, max := getBounds(cubes)
	return countReachableSurfaces(cubePositions, min, max)
}
