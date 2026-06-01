package world

import "fmt"

type Position struct {
	X int
	Y int
}

func (this Position) plus(other Position) Position {
	return Position{X: this.X + other.X, Y: this.Y + other.Y}
}

func (pos Position) String() string {
	return fmt.Sprintf("Position{X=%d, Y=%d)", pos.X, pos.Y)
}
