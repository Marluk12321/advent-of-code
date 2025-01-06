package mapping

import "fmt"

type FacingDirection int

const (
	FACING_RIGHT FacingDirection = 0
	FACING_DOWN  FacingDirection = 1
	FACING_LEFT  FacingDirection = 2
	FACING_UP    FacingDirection = 3
)

func (direction FacingDirection) Left() FacingDirection {
	switch direction {
	case FACING_RIGHT:
		return FACING_UP
	case FACING_DOWN:
		return FACING_RIGHT
	case FACING_LEFT:
		return FACING_DOWN
	case FACING_UP:
		return FACING_LEFT
	default:
		panic(direction)
	}
}

func (direction FacingDirection) Right() FacingDirection {
	switch direction {
	case FACING_RIGHT:
		return FACING_DOWN
	case FACING_DOWN:
		return FACING_LEFT
	case FACING_LEFT:
		return FACING_UP
	case FACING_UP:
		return FACING_RIGHT
	default:
		panic(direction)
	}
}

func (direction FacingDirection) Opposite() FacingDirection {
	switch direction {
	case FACING_RIGHT:
		return FACING_LEFT
	case FACING_DOWN:
		return FACING_UP
	case FACING_LEFT:
		return FACING_RIGHT
	case FACING_UP:
		return FACING_DOWN
	default:
		panic(direction)
	}
}

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

type Position struct {
	Row int
	Col int
}

func (position Position) String() string {
	return fmt.Sprint("Position(row:", position.Row, ", col:", position.Col, ")")
}

func (position Position) Copy() Position {
	return Position{
		Row: position.Row,
		Col: position.Col,
	}
}

type State struct {
	Position Position
	Facing   FacingDirection
}

func (state State) String() string {
	return fmt.Sprint("State(position:", state.Position, ", facing:", state.Facing, ")")
}
