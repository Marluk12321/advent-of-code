package instructions

import (
	"2022/22/part_2/mapping"
	"fmt"
)

type TurnDirection rune

const (
	RIGHT TurnDirection = 'R'
	LEFT  TurnDirection = 'L'
	BACK  TurnDirection = 'B'
)

func (turnDirection TurnDirection) String() string {
	switch turnDirection {
	case RIGHT:
		return "Right"
	case LEFT:
		return "Left"
	case BACK:
		return "Back"
	default:
		panic(turnDirection)
	}
}

type TurnInstruction struct {
	turnDirection TurnDirection
}

func (instruction TurnInstruction) String() string {
	return fmt.Sprintf("Turn %s", instruction.turnDirection.String())
}

func (instruction TurnInstruction) Apply(cube mapping.Cube, state mapping.State) mapping.State {
	fmt.Println("Turning", state, instruction.turnDirection)
	nextFacing := state.Facing
	switch instruction.turnDirection {
	case LEFT:
		nextFacing = nextFacing.Left()
	case RIGHT:
		nextFacing = nextFacing.Right()
	case BACK:
		nextFacing = nextFacing.Opposite()
	default:
		panic(instruction.turnDirection)
	}
	fmt.Println("Facing:", state.Facing, "->", nextFacing)
	return mapping.State{
		Face: state.Face,
		FacePosition: state.FacePosition.Copy(),
		Facing:   nextFacing,
	}
}
