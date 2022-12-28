package rope

type Position struct {
	X int
	Y int
}

func add(position1 *Position, position2 *Position) Position {
	return Position{
		X: position1.X + position2.X,
		Y: position1.Y + position2.Y,
	}
}

func sub(position1 *Position, position2 *Position) Position {
	return Position{
		X: position1.X - position2.X,
		Y: position1.Y - position2.Y,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (position *Position) Size() int {
	xSize := abs(position.X)
	ySize := abs(position.Y)
	if xSize > ySize {
		return xSize
	} else {
		return ySize
	}
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type RopeSegment struct {
	Position Position
}

func (segment *RopeSegment) move(direction Direction) {
	switch direction {
	case UP:
		segment.Position.Y++
	case DOWN:
		segment.Position.Y--
	case LEFT:
		segment.Position.X--
	case RIGHT:
		segment.Position.X++
	}
}

func clip(x int) int {
	if x < -1 {
		return -1
	} else if x > 1 {
		return 1
	} else {
		return x
	}
}

func toStep(position *Position) Position {
	return Position{
		X: clip(position.X),
		Y: clip(position.Y),
	}
}

func (segment *RopeSegment) follow(other *RopeSegment) {
	diff := sub(&other.Position, &segment.Position)
	if diff.Size() > 1 {
		step := toStep(&diff)
		segment.Position = add(&segment.Position, &step)
	}
}

type Rope struct {
	Head RopeSegment
	Tail RopeSegment
}

func (rope *Rope) Move(direction Direction) {
	rope.Head.move(direction)
	rope.Tail.follow(&rope.Head)
}
