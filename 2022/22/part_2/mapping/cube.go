package mapping

import (
	"2022/22/part_2/util"
)

type FaceOnBoard struct {
	Face     Face
	Position Position
	Edges    map[FacingDirection]*Edge
}

func makeFaceOnBoard(face Face, position Position) FaceOnBoard {
	return FaceOnBoard{
		Face:     face,
		Position: position,
		Edges:    make(map[FacingDirection]*Edge, 4),
	}
}

type Cube struct {
	Board Board
	Size  int
	Faces map[Face]*FaceOnBoard
}

func (cube Cube) GetTile(face Face, facePosition Position) Tile {
	faceOnBoard := cube.Faces[face]
	absolutePosition := faceOnBoard.Position.Add(facePosition)
	rowOnBoard := cube.Board[absolutePosition.Row]
	tileIndex := absolutePosition.Col - rowOnBoard.Offset
	return rowOnBoard.Tiles[tileIndex]
}

func findFacePositions(board Board, faceSize int) map[Position]bool {
	facePositions := make(map[Position]bool, 6)
	for row := 0; row < len(board); row += faceSize {
		boardRow := board[row]
		for col := 0; col < len(boardRow.Tiles); col += faceSize {
			facePosition := Position{Row: row, Col: col + boardRow.Offset}
			facePositions[facePosition] = true
			util.Debug("Face position:", facePosition)
		}
	}
	return facePositions
}

func buildFrontFace(board Board) FaceOnBoard {
	frontPosition := Position{
		Row: 0,
		Col: board[0].Offset,
	}
	return makeFaceOnBoard(FRONT_FACE, frontPosition)
}

func buildNeighborOffsets(size int) map[FacingDirection]Position {
	neighborOffsets := make(map[FacingDirection]Position, 4)
	for _, direction := range FacingDirections {
		neighborOffsets[direction] = UnitSteps[direction].Scale(size)
	}
	return neighborOffsets
}

func calcBoardToCubeDiff(faceOnBoard FaceOnBoard) FacingDirectionDifference {
	for edgeDirection, edge := range faceOnBoard.Edges {
		oppositeFace := edge.OppositeFace[faceOnBoard.Face]
		expectedDirection := faceOnBoard.Face.DirectionOf(oppositeFace)
		return edgeDirection.DiffTo(expectedDirection)
	}
	return DIFFERENCE_NONE
}

func calcCubeToBoardDiff(faceOnBoard FaceOnBoard) FacingDirectionDifference {
	return calcBoardToCubeDiff(faceOnBoard).Opposite()
}

func boardToCubeDirection(faceOnBoard *FaceOnBoard, boardDirection FacingDirection) FacingDirection {
	cubeDirection := boardDirection
	boardToCubeDiff := calcBoardToCubeDiff(*faceOnBoard)
	if boardToCubeDiff != DIFFERENCE_NONE {
		cubeDirection = boardDirection.Add(boardToCubeDiff)
		util.Debug(boardDirection, "becomes", cubeDirection)
	}
	return cubeDirection
}

func connectFaces(
		faceOnBoard1, faceOnBoard2 *FaceOnBoard,
		faceSize int,
		edgeDirection1, edgeDirection2 FacingDirection,
	) {
	edge := makeEdge(faceSize)
	faceOnBoard1.Edges[edgeDirection1] = &edge
	faceOnBoard2.Edges[edgeDirection2] = &edge
	edge.FaceDirection[faceOnBoard1.Face] = edgeDirection1.Opposite()
	edge.FaceDirection[faceOnBoard2.Face] = edgeDirection2.Opposite()
	edge.OppositeFace[faceOnBoard1.Face] = faceOnBoard2.Face
	edge.OppositeFace[faceOnBoard2.Face] = faceOnBoard1.Face
	util.Debug("Connected", faceOnBoard1, edgeDirection1, "<->", edgeDirection2, faceOnBoard2)
}

func connectAdjacentFaces(facePositions map[Position]bool, faceSize int, startingFace *FaceOnBoard) []*FaceOnBoard {
	adjacentlyConnectedFaces := []*FaceOnBoard{startingFace}
	queue := []*FaceOnBoard{startingFace}
	queueIndex := 0
	neighborOffsets := buildNeighborOffsets(faceSize)

	for queueIndex < len(queue) {
		faceOnBoard := queue[queueIndex]
		queueIndex++
		for direction, neighborOffset := range neighborOffsets {
			_, edgeExists := faceOnBoard.Edges[direction]
			if edgeExists {
				util.Debug("Edge exists", faceOnBoard.Face, direction)
				continue
			}
			neighborPosition := faceOnBoard.Position.Add(neighborOffset)
			if facePositions[neighborPosition] {
				util.Debug("Position exists", faceOnBoard.Face, direction, neighborPosition)
				cubeDirection := boardToCubeDirection(faceOnBoard, direction)
				neighborFace := faceOnBoard.Face.GetNeighbor(cubeDirection)
				neighborFaceOnBoard := makeFaceOnBoard(neighborFace, neighborPosition)
				connectFaces(faceOnBoard, &neighborFaceOnBoard, faceSize, direction, direction.Opposite())
				adjacentlyConnectedFaces = append(adjacentlyConnectedFaces, &neighborFaceOnBoard)
				queue = append(queue, &neighborFaceOnBoard)
			} else {
				util.Debug("Position empty", faceOnBoard.Face, direction, neighborPosition)
			}
		}
	}
	return adjacentlyConnectedFaces
}

func connectMissingEdges(boardFaces map[Face]*FaceOnBoard, faceSize int) {
	for face, faceOnBoard := range boardFaces {
		boardToCubeDiff := calcBoardToCubeDiff(*faceOnBoard)
		for _, direction := range FacingDirections {
			_, edgeExists := faceOnBoard.Edges[direction]
			if edgeExists {
				continue
			}
			cubeDirection := direction.Add(boardToCubeDiff)
			neighborFace := face.GetNeighbor(cubeDirection)
			neighborFaceOnBoard := boardFaces[neighborFace]
			neighborCubeToBoardDiff := calcCubeToBoardDiff(*neighborFaceOnBoard)
			expectedReturnDirection := neighborFace.DirectionOf(face)
			returnDirection := expectedReturnDirection.Add(neighborCubeToBoardDiff)
			connectFaces(faceOnBoard, neighborFaceOnBoard, faceSize, direction, returnDirection)
		}
	}
}

func mapToFaces(board Board, faceSize int) map[Face]*FaceOnBoard {
	facePositions := findFacePositions(board, faceSize)
	frontFace := buildFrontFace(board)
	adjacentlyConnectedFaces := connectAdjacentFaces(facePositions, faceSize, &frontFace)
	boardFaces := make(map[Face]*FaceOnBoard, 6)
	for _, faceOnBoard := range adjacentlyConnectedFaces {
		boardFaces[faceOnBoard.Face] = faceOnBoard
	}
	connectMissingEdges(boardFaces, faceSize)
	return boardFaces
}

func BuildCube(board Board, size int) Cube {
	return Cube{
		Board: board,
		Size:  size,
		Faces: mapToFaces(board, size),
	}
}
