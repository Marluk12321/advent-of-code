package mapping

import "fmt"

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

func mapToFaces(board Board, faceSize int) map[Face]*FaceOnBoard {
	facePositions := make(map[Position]bool, 6)
	for row := 0; row < len(board); row += faceSize {
		boardRow := board[row]
		for col := 0; col < len(boardRow.Tiles); col += faceSize {
			facePosition := Position{Row: row, Col: col + boardRow.Offset}
			facePositions[facePosition] = true
			fmt.Println("Face position:", facePosition)
		}
	}

	frontPosition := Position{
		Row: 0,
		Col: board[0].Offset,
	}
	frontFace := makeFaceOnBoard(FRONT_FACE, frontPosition)
	boardFaces := map[Face]*FaceOnBoard{
		FRONT_FACE: &frontFace,
	}

	queue := []*FaceOnBoard{&frontFace}
	queueIndex := 0
	neighborOffsets := make(map[FacingDirection]Position, 4)
	for _, direction := range FacingDirections {
		neighborOffsets[direction] = UnitSteps[direction].Scale(faceSize)
	}
	for queueIndex < len(queue) {
		faceOnBoard := queue[queueIndex]
		queueIndex++
		for direction, neighborOffset := range neighborOffsets {
			_, edgeExists := faceOnBoard.Edges[direction]
			if edgeExists {
				fmt.Println("Edge exists", faceOnBoard.Face, direction)
				continue
			}
			neighborPosition := faceOnBoard.Position.Add(neighborOffset)
			if facePositions[neighborPosition] {
				fmt.Println("Position exists", faceOnBoard.Face, direction, neighborPosition)
				expectedNeighborDirection := direction
				for edgeDirection, edge := range faceOnBoard.Edges {
					oppositeFace := edge.OppositeFace[faceOnBoard.Face]
					expectedDirection := faceOnBoard.Face.DirectionOf(oppositeFace)
					if edgeDirection != expectedDirection {
						fmt.Println(oppositeFace, "was supposed to be", expectedDirection, "but was", edgeDirection)
						directionDifference := edgeDirection.DiffTo(expectedDirection)
						expectedNeighborDirection = expectedNeighborDirection.Add(directionDifference)
						fmt.Println(direction, "becomes", expectedNeighborDirection)
					}
					break
				}
				neighborFace := faceOnBoard.Face.GetNeighbor(expectedNeighborDirection)
				neighborFaceOnBoard := makeFaceOnBoard(neighborFace, neighborPosition)
				//connectNeighbors(faceOnBoard, direction, neighborFaceOnBoard)
				edge := makeEdge(faceSize)
				faceOnBoard.Edges[direction] = &edge
				neighborFaceOnBoard.Edges[direction.Opposite()] = &edge
				edge.FaceDirection[faceOnBoard.Face] = direction.Opposite()
				edge.FaceDirection[neighborFaceOnBoard.Face] = direction
				edge.OppositeFace[faceOnBoard.Face] = neighborFaceOnBoard.Face
				edge.OppositeFace[neighborFaceOnBoard.Face] = faceOnBoard.Face
				fmt.Println("Connected", faceOnBoard.Face, direction, neighborFace)
				boardFaces[neighborFace] = &neighborFaceOnBoard
				queue = append(queue, &neighborFaceOnBoard)
			} else {
				fmt.Println("Position empty", faceOnBoard.Face, direction, neighborPosition)
			}
		}
	}

	for face, faceOnBoard := range boardFaces {
		directionDifference := DIFFERENCE_NONE
		for edgeDirection, edge := range faceOnBoard.Edges {
			oppositeFace := edge.OppositeFace[face]
			expectedDirection := face.DirectionOf(oppositeFace)
			directionDifference = edgeDirection.DiffTo(expectedDirection)
			break
		}
		for _, direction := range FacingDirections {
			_, edgeExists := faceOnBoard.Edges[direction]
			if edgeExists {
				continue
			}
			expectedNeighborDirection := direction.Add(directionDifference)
			neighborFace := face.GetNeighbor(expectedNeighborDirection)
			neighborFaceOnBoard := boardFaces[neighborFace]

			neighborDirectionDifference := DIFFERENCE_NONE
			for edgeDirection, edge := range neighborFaceOnBoard.Edges {
				oppositeFace := edge.OppositeFace[neighborFace]
				expectedDirection := neighborFace.DirectionOf(oppositeFace)
				neighborDirectionDifference = expectedDirection.DiffTo(edgeDirection)
				break
			}
			expectedReturnDirection := neighborFace.DirectionOf(face)
			returnDirection := expectedReturnDirection.Add(neighborDirectionDifference)

			edge := makeEdge(faceSize)
			faceOnBoard.Edges[direction] = &edge
			neighborFaceOnBoard.Edges[returnDirection] = &edge
			edge.FaceDirection[face] = direction.Opposite()
			edge.FaceDirection[neighborFace] = returnDirection.Opposite()
			edge.OppositeFace[face] = neighborFace
			edge.OppositeFace[neighborFace] = face
			fmt.Println("Connected", face, direction, "<->", returnDirection, neighborFace)
		}
	}

	return boardFaces
}

func BuildCube(board Board, size int) Cube {
	return Cube{
		Board: board,
		Size:  size,
		Faces: mapToFaces(board, size),
	}
}
