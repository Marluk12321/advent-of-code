package world

type Direction int
const(
	UP Direction = iota
	DOWN Direction = iota
	LEFT Direction = iota
	RIGHT Direction = iota
)

func (direction Direction) String() string {
	switch direction {
	case UP:
		return "^"
	case DOWN:
		return "v"
	case LEFT:
		return "<"
	case RIGHT:
		return ">"
	default:
		panic(direction)
	}
}

type Blizzard struct {
	Row, Col int
	Direction Direction
}
