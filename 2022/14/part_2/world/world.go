package world

type Position struct {
	X int
	Y int
}

type World struct {
	RockFormations []RockFormation
	SandPositions  []Position
	FloorY         int
}

func MakeWorld(rockFormations []RockFormation) World {
	maxY := rockFormations[0].Max.Y
	for i := range rockFormations[1:] {
		formation := &rockFormations[i+1]
		if formation.Max.Y > maxY {
			maxY = formation.Max.Y
		}
	}
	return World{
		RockFormations: rockFormations,
		FloorY:         maxY + 2,
	}
}
