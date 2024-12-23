package board

import "fmt"

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

func (instruction TurnInstruction) Apply(boardMap Map, state State) State {
	nextState := State{
		Row:    state.Row,
		Col:    state.Col,
		Facing: state.Facing,
	}
	fmt.Println("Turning", state, instruction.turnDirection)
	switch state.Facing {
	case FACING_RIGHT:
		switch instruction.turnDirection {
		case LEFT:
			nextState.Facing = FACING_UP
		case RIGHT:
			nextState.Facing = FACING_DOWN
		default:
			panic(instruction.turnDirection)
		}
	case FACING_DOWN:
		switch instruction.turnDirection {
		case LEFT:
			nextState.Facing = FACING_RIGHT
		case RIGHT:
			nextState.Facing = FACING_LEFT
		default:
			panic(instruction.turnDirection)
		}
	case FACING_LEFT:
		switch instruction.turnDirection {
		case LEFT:
			nextState.Facing = FACING_DOWN
		case RIGHT:
			nextState.Facing = FACING_UP
		default:
			panic(instruction.turnDirection)
		}
	case FACING_UP:
		switch instruction.turnDirection {
		case LEFT:
			nextState.Facing = FACING_LEFT
		case RIGHT:
			nextState.Facing = FACING_RIGHT
		default:
			panic(instruction.turnDirection)
		}
	default:
		panic(state.Facing)
	}
	return nextState
}
