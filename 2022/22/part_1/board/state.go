package board

import "fmt"

type FacingDirection int

const (
	FACING_RIGHT FacingDirection = 0
	FACING_DOWN  FacingDirection = 1
	FACING_LEFT  FacingDirection = 2
	FACING_UP    FacingDirection = 3
)

func (direction FacingDirection) String() string {
	switch direction {
	case FACING_RIGHT:
		return "Right"
	case FACING_DOWN:
		return "Down"
	case FACING_LEFT:
		return "Left"
	case FACING_UP:
		return "Up"
	default:
		panic(direction)
	}
}

type State struct {
	Row    int
	Col    int
	Facing FacingDirection
}

func (state State) String() string {
	return fmt.Sprint("State(row:", state.Row, ", col:", state.Col, ", facing:", state.Facing, ")")
}
