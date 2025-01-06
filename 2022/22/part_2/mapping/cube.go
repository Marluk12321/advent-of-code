package mapping

type Face int

const (
	FRONT Face = iota
	BACK
	TOP
	BOTTOM
	LEFT
	RIGHT
)

func (face Face) Opposite() Face {
	switch face {
	case FRONT:
		return BACK
	case BACK:
		return FRONT
	case TOP:
		return BOTTOM
	case BOTTOM:
		return TOP
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	default:
		panic(face)
	}
}

func (face Face) getAdjacent(direction FacingDirection) Face {
	switch direction {
	case FACING_UP:
		switch face {
		case TOP:
			return BACK
		case BOTTOM:
			return FRONT
		case BACK:
			return BOTTOM
		default:
			return TOP
		}
	case FACING_DOWN:
		return face.getAdjacent(FACING_UP).Opposite()
	case FACING_LEFT:
		switch face {
		case LEFT:
			return BACK
		case RIGHT:
			return FRONT
		case BACK:
			return RIGHT
		default:
			return LEFT
		}
	case FACING_RIGHT:
		return face.getAdjacent(FACING_LEFT).Opposite()
	default:
		panic(direction)
	}
}

type CubeDirection int

const (
	CUBE_UP CubeDirection = iota
	CUBE_DOWN
	CUBE_LEFT
	CUBE_RIGHT
	CUBE_FRONT
	CUBE_BACK
)

type Orientation struct {
	Front CubeDirection
	Top   CubeDirection
}

func (orientation Orientation) Rotate(direction FacingDirection) Orientation {
	rotatedOrientation := Orientation{
		Front: orientation.Front,
		Top:   orientation.Top,
	}
	switch direction {
	case FACING_UP:

	default:
		panic(direction)
	}
	return rotatedOrientation
}

type FaceOnBoard struct {
	Face        Face
	Position    Position
	Orientation Orientation
}

type Cube struct {
	Board Board
	Size  int
}

func makeFrontBoardFace(board Board) FaceOnBoard {
	return FaceOnBoard{
		Face:        FRONT,
		Position:    Position{Row: 0, Col: board[0].Offset},
		Orientation: Orientation{Front: CUBE_FRONT, Top: CUBE_UP},
	}
}

func normalizeDirection(direction FacingDirection, orientation Orientation) FacingDirection {
	
}

func determineAdjacentFace(boardFace FaceOnBoard, direction FacingDirection) Face {
	normalizedDirection := normalizeDirection(direction, boardFace.Orientation)
	return boardFace.Face.getAdjacent(normalizedDirection)
}

func makeAdjecant(boardFace FaceOnBoard, direction FacingDirection, position Position) FaceOnBoard {
	adjecantFace := determineAdjacentFace(boardFace, direction)
	adjacentOrientation := boardFace.Orientation.Rotate(direction)
	return FaceOnBoard{
		Face:        adjecantFace,
		Position:    position,
		Orientation: adjacentOrientation,
	}
}

func mapToFaces(board Board, faceSize int) map[Face]FaceOnBoard {
	boardFaces := make(map[Face]FaceOnBoard, 6)
	boardFaces[FRONT] = makeFrontBoardFace(board)
	leftBoardFace := boardFaces[FRONT]
	for row := 0; row < len(board); row += faceSize {
		boardRow := board[row]
		for col := boardRow.Offset + faceSize; col < len(boardRow.Tiles); col += faceSize {
			rightPosition := Position{Row: row, Col: col}
			rightBoardFace := makeAdjecant(leftBoardFace, FACING_RIGHT, rightPosition)
			boardFaces[rightBoardFace.Face] = rightBoardFace
			leftBoardFace = rightBoardFace
		}
	}

	return boardFaces
}

func BuildCube(board Board, size int) Cube {
	return Cube{
		Board: board,
		Size:  size,
	}
}
