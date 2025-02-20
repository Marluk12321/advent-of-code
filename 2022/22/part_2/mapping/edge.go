package mapping

var edgeStartCorner = map[FacingDirection]Position{
	FACING_UP:    {0, 0},
	FACING_LEFT:  {0, 0},
	FACING_RIGHT: {0, 1},
	FACING_DOWN:  {1, 0},
}

var haveOppositeStartCorner = map[[2]FacingDirection]bool{
	{FACING_UP, FACING_UP}:  true,
	{FACING_UP, FACING_RIGHT}:  true,
	{FACING_DOWN, FACING_DOWN}: true,
	{FACING_DOWN, FACING_LEFT}: true,
	{FACING_RIGHT, FACING_RIGHT}:  true,
	{FACING_RIGHT, FACING_UP}:  true,
	{FACING_LEFT, FACING_LEFT}: true,
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

func calcDistanceFromEdgeStart(edgeDirection FacingDirection, edgeSize int, position Position) int {
	edgeStart := edgeStartCorner[edgeDirection].Scale(edgeSize - 1)
	edgeStartToPosition := position.Sub(edgeStart)
	if edgeStartToPosition.Row != 0 && edgeStartToPosition.Col != 0 {
		panic(position)
	}
	return edgeStartToPosition.ChebyshevSize()
}

func makeOppositePosition(edgeDirection, oppositeEdgeDirection FacingDirection, edgeSize int, position Position) Position {
	distanceFromEdgeStart := calcDistanceFromEdgeStart(edgeDirection, edgeSize, position)
	if haveOppositeStartCorner[[2]FacingDirection{edgeDirection, oppositeEdgeDirection}] {
		distanceFromEdgeStart = edgeSize - 1 - distanceFromEdgeStart
	}
	oppositeEdgeStart := edgeStartCorner[oppositeEdgeDirection].Scale(edgeSize - 1)
	oppositePositionOnEdge := edgeStep[oppositeEdgeDirection].Scale(distanceFromEdgeStart)
	return oppositeEdgeStart.Add(oppositePositionOnEdge)
}

func (edge Edge) MakeOppositeState(state State) State {
	oppositeFace := edge.OppositeFace[state.Face]
	oppositeFacing := edge.FaceDirection[oppositeFace]
	edgeDirection := state.Facing
	oppositeEdgeDirection := oppositeFacing.Opposite()
	oppositePosition := makeOppositePosition(edgeDirection, oppositeEdgeDirection, edge.Size, state.FacePosition)
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
