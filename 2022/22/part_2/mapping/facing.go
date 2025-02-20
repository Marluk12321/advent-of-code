package mapping

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

type FacingDirectionDifference int

const (
	DIFFERENCE_NONE FacingDirectionDifference = iota
	DIFFERENCE_LEFT
	DIFFERENCE_RIGHT
	DIFFERENCE_OPPOSITE
)

func (direction FacingDirection) DiffTo(other FacingDirection) FacingDirectionDifference {
	if direction == other {
		return DIFFERENCE_NONE
	}
	if direction.Left() == other {
		return DIFFERENCE_LEFT
	}
	if direction.Right() == other {
		return DIFFERENCE_RIGHT
	}
	return DIFFERENCE_OPPOSITE
}

func (direction FacingDirection) Add(difference FacingDirectionDifference) FacingDirection {
	switch difference {
	case DIFFERENCE_NONE:
		return direction
	case DIFFERENCE_LEFT:
		return direction.Left()
	case DIFFERENCE_RIGHT:
		return direction.Right()
	case DIFFERENCE_OPPOSITE:
		return direction.Opposite()
	default:
		panic(difference)
	}
}

var FacingDirections = [4]FacingDirection{
	FACING_UP,
	FACING_DOWN,
	FACING_LEFT,
	FACING_RIGHT,
}
var UnitSteps = map[FacingDirection]Position{
	FACING_UP:    {Row: -1, Col: 0},
	FACING_DOWN:  {Row: 1, Col: 0},
	FACING_LEFT:  {Row: 0, Col: -1},
	FACING_RIGHT: {Row: 0, Col: 1},
}
