package mapping

import "fmt"

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

func (position Position) Add(other Position) Position {
	return Position{
		Row: position.Row + other.Row,
		Col: position.Col + other.Col,
	}
}

func (position Position) Sub(other Position) Position {
	return Position{
		Row: position.Row - other.Row,
		Col: position.Col - other.Col,
	}
}

func (position Position) Scale(amount int) Position {
	return Position{
		Row: position.Row * amount,
		Col: position.Col * amount,
	}
}

func (position Position) ChebyshevSize() int {
	row := position.Row
	if row < 0 {
		row = -row
	}
	col := position.Col
	if col < 0 {
		col = -col
	}
	if row > col {
		return row
	} else {
		return col
	}
}
