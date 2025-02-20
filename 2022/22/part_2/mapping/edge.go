package mapping

var edgeStartCorner = map[FacingDirection]Position{
	FACING_UP:    {0, 0},
	FACING_LEFT:  {0, 0},
	FACING_RIGHT: {0, 1},
	FACING_DOWN:  {1, 0},
}

var haveOppositeStartCorner = map[[2]FacingDirection]bool{
	{FACING_UP, FACING_RIGHT}:  true,
	{FACING_DOWN, FACING_LEFT}: true,
	{FACING_RIGHT, FACING_UP}:  true,
	{FACING_LEFT, FACING_DOWN}: true,
}

var edgeStep = map[FacingDirection]Position{
	FACING_UP:    UnitSteps[FACING_RIGHT],
	FACING_LEFT:  UnitSteps[FACING_DOWN],
	FACING_RIGHT: UnitSteps[FACING_DOWN],
	FACING_DOWN:  UnitSteps[FACING_RIGHT],
}

type Edge struct {
	Size          int
	OppositeFace  map[Face]Face
	FaceDirection map[Face]FacingDirection
}

func (edge Edge) MakeOppositeState(state State) State {
	edgeDirection := state.Facing
	edgeStartPosition := edgeStartCorner[edgeDirection].Scale(edge.Size - 1)
	positionOnEdge := state.FacePosition.Sub(edgeStartPosition)
	distanceFromEdgeStart := positionOnEdge.ChebyshevSize()

	oppositeFace := edge.OppositeFace[state.Face]
	oppositeFacing := edge.FaceDirection[oppositeFace]
	oppositeEdgeDirection := oppositeFacing.Opposite()
	if edgeDirection == oppositeEdgeDirection || haveOppositeStartCorner[[2]FacingDirection{edgeDirection, oppositeEdgeDirection}] {
		distanceFromEdgeStart = edge.Size - 1 - distanceFromEdgeStart
	}
	oppositeEdgeStartPosition := edgeStartCorner[oppositeEdgeDirection].Scale(edge.Size - 1)
	oppositePositionOnEdge := edgeStep[oppositeEdgeDirection].Scale(distanceFromEdgeStart)
	oppositePosition := oppositeEdgeStartPosition.Add(oppositePositionOnEdge)
	return State{
		Face:         oppositeFace,
		FacePosition: oppositePosition,
		Facing:       oppositeFacing,
	}
}

func makeEdge(size int) Edge {
	return Edge{
		Size:          size,
		OppositeFace:  make(map[Face]Face, 2),
		FaceDirection: make(map[Face]FacingDirection, 2),
	}
}
