package instructions

import (
	"2022/22/part_2/mapping"
	"2022/22/part_2/util"
	"fmt"
)

type TurnDirection rune

const (
	RIGHT TurnDirection = 'R'
	LEFT  TurnDirection = 'L'
)

func (turnDirection TurnDirection) String() string {
	switch turnDirection {
	case RIGHT:
		return "Right"
	case LEFT:
		return "Left"
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
	util.Debug(state, instruction)
	nextFacing := state.Facing
	switch instruction.turnDirection {
	case LEFT:
		nextFacing = nextFacing.Left()
	case RIGHT:
		nextFacing = nextFacing.Right()
	default:
		panic(instruction.turnDirection)
	}
	util.Debug("Facing:", state.Facing, "->", nextFacing)
	return mapping.State{
		Face: state.Face,
		FacePosition: state.FacePosition.Copy(),
		Facing:   nextFacing,
	}
}
