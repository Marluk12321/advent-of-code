package world

type Direction int

const (
	N    Direction = iota
	S    Direction = iota
	W    Direction = iota
	E    Direction = iota
	NW   Direction = iota
	NE   Direction = iota
	SW   Direction = iota
	SE   Direction = iota
	NONE Direction = iota
)

func (direction Direction) String() string {
	switch direction {
	case N:
		return "N"
	case S:
		return "S"
	case W:
		return "W"
	case E:
		return "E"
	case NW:
		return "NW"
	case NE:
		return "NE"
	case SW:
		return "SW"
	case SE:
		return "SE"
	case NONE:
		return "NONE"
	default:
		panic(direction)
	}
}

func (direction Direction) Opposite() Direction {
	switch direction {
	case N:
		return S
	case S:
		return N
	case W:
		return E
	case E:
		return W
	case NW:
		return SE
	case NE:
		return SW
	case SW:
		return NE
	case SE:
		return NW
	case NONE:
		return NONE
	default:
		panic(direction)
	}
}

func (direction Direction) DecomposeToOrthogonals() []Direction {
	switch direction {
	case N:
		return []Direction{N}
	case S:
		return []Direction{S}
	case W:
		return []Direction{W}
	case E:
		return []Direction{E}
	case NW:
		return []Direction{N, W}
	case NE:
		return []Direction{N, E}
	case SW:
		return []Direction{S, W}
	case SE:
		return []Direction{S, E}
	case NONE:
		return []Direction{NONE}
	default:
		panic(direction)
	}
}
