package calculation

import "2022/18/part_1/objects"

func getNeighbors(cube objects.Cube) [6]objects.Position {
	return [6]objects.Position{
		{X: cube.X - 1, Y: cube.Y, Z: cube.Z},
		{X: cube.X + 1, Y: cube.Y, Z: cube.Z},
		{X: cube.X, Y: cube.Y - 1, Z: cube.Z},
		{X: cube.X, Y: cube.Y + 1, Z: cube.Z},
		{X: cube.X, Y: cube.Y, Z: cube.Z - 1},
		{X: cube.X, Y: cube.Y, Z: cube.Z + 1},
	}
}

func CalcTotalSurface(cubes []objects.Cube) int {
	exposedSides := len(cubes) * 6
	neighboringCubes := map[objects.Position]int{}
	for _, cube := range cubes {
		exposedSides -= 2 * neighboringCubes[objects.Position(cube)]
		for _, neighboringPosition := range getNeighbors(cube) {
			neighboringCubes[neighboringPosition]++
		}
	}
	return exposedSides
}
