package instructions

import (
	"2022/22/part_2/mapping"
	"fmt"
)

type MoveInstruction struct {
	moveBy int
}

func (instruction MoveInstruction) String() string {
	return fmt.Sprintf("Move by %d", instruction.moveBy)
}

func findWrappingRow(cube mapping.Cube, state mapping.State, step int) int {
	row := state.Position.Row
	for {
		nextRow := (row + step) % len(cube)
		if nextRow < 0 {
			nextRow += len(cube)
		}

		nextBoardRow := cube[nextRow]
		if !nextBoardRow.IsWithinBounds(state.Position.Col) {
			break
		}
		row = nextRow
	}
	return row
}

func moveHorizontally(cube mapping.Cube, state mapping.State, by, step int) mapping.State {
	boardRow := cube[state.Position.Row]
	relativeCol := state.Position.Col - boardRow.Offset
moveLoop:
	for i := 0; i < by; i++ {
		nextCol := (relativeCol + step) % len(boardRow.Tiles)
		if nextCol < 0 {
			nextCol += len(boardRow.Tiles)
		}

		tile := boardRow.Tiles[nextCol]
		switch tile {
		case mapping.WALL:
			break moveLoop
		case mapping.OPEN:
			fmt.Println("col:", relativeCol, "->", nextCol)
			relativeCol = nextCol
		default:
			panic(tile)
		}
	}
	return mapping.State{
		Position: mapping.Position{
			Row:    state.Position.Row,
			Col:    relativeCol + boardRow.Offset,
		},
		Facing: state.Facing,
	}
}

func moveVertically(cube mapping.Cube, state mapping.State, by, step int) mapping.State {
	row := state.Position.Row
moveLoopDown:
	for i := 0; i < by; i++ {
		nextRow := (row + step) % len(cube)
		if nextRow < 0 {
			nextRow += len(cube)
		}

		boardRow := cube[nextRow]
		if !boardRow.IsWithinBounds(state.Position.Col) {
			nextRow = findWrappingRow(cube, state, -step)
			if nextRow == row {
				break
			}
			boardRow = cube[nextRow]
		}

		tile := boardRow.Tiles[state.Position.Col-boardRow.Offset]
		switch tile {
		case mapping.WALL:
			break moveLoopDown
		case mapping.OPEN:
			fmt.Println("row:", row, "->", nextRow)
			row = nextRow
		default:
			panic(tile)
		}
	}
	return mapping.State{
		Position: mapping.Position{
			Row:    row,
			Col:    state.Position.Col,
		},
		Facing: state.Facing,
	}
}

func (instruction MoveInstruction) Apply(cube mapping.Cube, state mapping.State) mapping.State {
	fmt.Println("Moving", state, "by", instruction.moveBy)
	switch state.Facing {
	case mapping.FACING_RIGHT:
		return moveHorizontally(cube, state, instruction.moveBy, 1)
	case mapping.FACING_DOWN:
		return moveVertically(cube, state, instruction.moveBy, 1)
	case mapping.FACING_LEFT:
		return moveHorizontally(cube, state, instruction.moveBy, -1)
	case mapping.FACING_UP:
		return moveVertically(cube, state, instruction.moveBy, -1)
	default:
		panic(state.Facing)
	}
}
