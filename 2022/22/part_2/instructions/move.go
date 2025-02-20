package instructions

import (
	"2022/22/part_2/mapping"
	"fmt"
)

type MoveInstruction struct {
	moveBy int
}

func (instruction MoveInstruction) String() string {
	return fmt.Sprintf("Move by %d", instruction.moveBy)
}

func isOutsideFace(position mapping.Position, faceSize int) bool {
	return position.Row < 0 || position.Col < 0 || position.Row >= faceSize || position.Col >= faceSize
}

func moveToNextState(cube mapping.Cube, currentState mapping.State) mapping.State {
	stepOffset := mapping.UnitSteps[currentState.Facing]
	nextPosition := currentState.FacePosition.Add(stepOffset)
	if isOutsideFace(nextPosition, cube.Size) {
		faceOnBoard := cube.Faces[currentState.Face]
		edge := faceOnBoard.Edges[currentState.Facing]
		return edge.MakeOppositeState(currentState)
	} else {
		return mapping.State{
			Face:         currentState.Face,
			FacePosition: nextPosition,
			Facing:       currentState.Facing,
		}
	}
}

func (instruction MoveInstruction) Apply(cube mapping.Cube, state mapping.State) mapping.State {
	fmt.Println("Moving", state, "by", instruction.moveBy)
	currentState := state
LOOP:
	for i := 0; i < instruction.moveBy; i++ {
		nextState := moveToNextState(cube, currentState)
		nextTile := cube.GetTile(nextState.Face, nextState.FacePosition)
		switch nextTile {
		case mapping.WALL:
			fmt.Println("Wall at", nextState.Face, nextState.FacePosition)
			break LOOP
		case mapping.OPEN:
			fmt.Println("state:", currentState, "->", nextState)
			currentState = nextState
		default:
			panic(nextTile)
		}
	}
	return currentState
}
