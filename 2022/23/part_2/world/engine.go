package world

import (
	"2022/23/part_2/util"
)

var directionSteps = map[Direction]Position{
	N:    {X: 0, Y: -1},
	S:    {X: 0, Y: 1},
	W:    {X: -1, Y: 0},
	E:    {X: 1, Y: 0},
	NW:   {X: -1, Y: -1},
	NE:   {X: 1, Y: -1},
	SW:   {X: -1, Y: 1},
	SE:   {X: 1, Y: 1},
	NONE: {X: 0, Y: 0},
}

var adjacentDirections = map[Direction][]Direction{
	N: {NW, N, NE},
	S: {SW, S, SE},
	W: {NW, W, SW},
	E: {NE, E, SE},
}

func FindStableState(elves []Position) int {
	moveDirectionOrder := []Direction{N, S, W, E}
	iteration := 0

	for true {
		util.Debug("\nStarting iteration", iteration)
		isPositionOccupied := make(map[Position]bool)
		isDirectionBlocked := make(map[Position]map[Direction]bool)
		for _, elf := range elves {
			isPositionOccupied[elf] = true
			isDirectionBlocked[elf] = make(map[Direction]bool)
		}
		nextElves := make([]Position, len(elves))

		for i, elf := range elves {
			firstFreeDirection := NONE
			allDirectionsClear := true

			for _, moveDirection := range moveDirectionOrder {
				if firstFreeDirection != NONE && !allDirectionsClear {
					util.Debug("Stopping search", elf, firstFreeDirection)
					break
				}
				util.Debug("Checking", elf, moveDirection)

				if isDirectionBlocked[elf][moveDirection] {
					util.Debug("Direction blocked", elf, moveDirection)
					allDirectionsClear = false
					continue
				}

				for _, adjacentDirection := range adjacentDirections[moveDirection] {
					adjacentStep := directionSteps[adjacentDirection]
					adjacentPosition := elf.plus(adjacentStep)
					util.Debug("Checking adjacent", adjacentPosition)
					if isPositionOccupied[adjacentPosition] {
						// if NE is occupied, N and E moves are blocked
						// S and W are also blocked for the neighbor
						for _, orthogonalDirection := range adjacentDirection.DecomposeToOrthogonals() {
							util.Debug("Marking blocked", elf, orthogonalDirection)
							isDirectionBlocked[elf][orthogonalDirection] = true
							util.Debug("Marking blocked", adjacentPosition, orthogonalDirection.Opposite())
							isDirectionBlocked[adjacentPosition][orthogonalDirection.Opposite()] = true
						}
						allDirectionsClear = false
						break
					}
				}

				if firstFreeDirection == NONE && !isDirectionBlocked[elf][moveDirection] {
					util.Debug("First free direction", elf, moveDirection)
					firstFreeDirection = moveDirection
				}
			}

			nextMoveDirection := NONE
			if !allDirectionsClear {
				nextMoveDirection = firstFreeDirection
			}
			util.Debug("Moving", elf, nextMoveDirection)
			nextMoveStep := directionSteps[nextMoveDirection]
			nextElves[i] = elf.plus(nextMoveStep)
			util.Debug("Next elf", elf, nextElves[i])
		}

		positionOccupants := make(map[Position]int)
		for _, nextElf := range nextElves {
			positionOccupants[nextElf] += 1
		}
		noElfMoved := true
		for i, nextElf := range nextElves {
			util.Debug("Occupants", nextElf, positionOccupants[nextElf])
			if positionOccupants[nextElf] == 1 && nextElf != elves[i] {
				util.Debug("Replacing", elves[i], nextElf)
				elves[i] = nextElf
				noElfMoved = false
			}
		}
		if noElfMoved {
			util.Debug("Stopping at iteration", iteration)
			return iteration
		}

		moveDirectionOrder = append(moveDirectionOrder[1:], moveDirectionOrder[0])
		iteration += 1
	}

	return -1
}
