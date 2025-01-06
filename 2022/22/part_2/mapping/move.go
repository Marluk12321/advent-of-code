package mapping

import "fmt"

type MoveInstruction struct {
	moveBy int
}

func (instruction MoveInstruction) String() string {
	return fmt.Sprintf("Move by %d", instruction.moveBy)
}

func findWrappingRow(boardMap Board, state State, step int) int {
	row := state.Row
	for {
		nextRow := (row + step) % len(boardMap)
		if nextRow < 0 {
			nextRow += len(boardMap)
		}

		nextBoardRow := boardMap[nextRow]
		if !nextBoardRow.IsWithinBounds(state.Col) {
			break
		}
		row = nextRow
	}
	return row
}

func moveHorizontally(boardMap Board, state State, by, step int) State {
	boardRow := boardMap[state.Row]
	relativeCol := state.Col - boardRow.Offset
moveLoop:
	for i := 0; i < by; i++ {
		nextCol := (relativeCol + step) % len(boardRow.Tiles)
		if nextCol < 0 {
			nextCol += len(boardRow.Tiles)
		}

		tile := boardRow.Tiles[nextCol]
		switch tile {
		case WALL:
			break moveLoop
		case OPEN:
			fmt.Println("col:", relativeCol, "->", nextCol)
			relativeCol = nextCol
		default:
			panic(tile)
		}
	}
	return State{
		Row:    state.Row,
		Col:    relativeCol + boardRow.Offset,
		Facing: state.Facing,
	}
}

func moveVertically(boardMap Board, state State, by, step int) State {
	row := state.Row
moveLoopDown:
	for i := 0; i < by; i++ {
		nextRow := (row + step) % len(boardMap)
		if nextRow < 0 {
			nextRow += len(boardMap)
		}

		boardRow := boardMap[nextRow]
		if !boardRow.IsWithinBounds(state.Col) {
			nextRow = findWrappingRow(boardMap, state, -step)
			if nextRow == row {
				break
			}
			boardRow = boardMap[nextRow]
		}

		tile := boardRow.Tiles[state.Col-boardRow.Offset]
		switch tile {
		case WALL:
			break moveLoopDown
		case OPEN:
			fmt.Println("row:", row, "->", nextRow)
			row = nextRow
		default:
			panic(tile)
		}
	}
	return State{
		Row:    row,
		Col:    state.Col,
		Facing: state.Facing,
	}
}

func (instruction MoveInstruction) Apply(boardMap Board, state State) State {
	fmt.Println("Moving", state, "by", instruction.moveBy)
	switch state.Facing {
	case FACING_RIGHT:
		return moveHorizontally(boardMap, state, instruction.moveBy, 1)
	case FACING_DOWN:
		return moveVertically(boardMap, state, instruction.moveBy, 1)
	case FACING_LEFT:
		return moveHorizontally(boardMap, state, instruction.moveBy, -1)
	case FACING_UP:
		return moveVertically(boardMap, state, instruction.moveBy, -1)
	default:
		panic(state.Facing)
	}
}
