package world

type Position struct {
	X int
	Y int
}

type World struct {
	RockFormations []RockFormation
	SandPositions  []Position
}

func MakeWorld(rockFormations []RockFormation) World {
	return World{
		RockFormations: rockFormations,
	}
}
