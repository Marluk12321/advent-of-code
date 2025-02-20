package mapping

type Face int

const (
	FRONT_FACE Face = iota
	BACK_FACE
	TOP_FACE
	BOTTOM_FACE
	LEFT_FACE
	RIGHT_FACE
)

func (face Face) String() string {
	switch face {
	case FRONT_FACE:
		return "FRONT"
	case BACK_FACE:
		return "BACK"
	case TOP_FACE:
		return "TOP"
	case BOTTOM_FACE:
		return "BOTTOM"
	case LEFT_FACE:
		return "LEFT"
	case RIGHT_FACE:
		return "RIGHT"
	default:
		panic(face)
	}
}

func (face Face) Opposite() Face {
	switch face {
	case FRONT_FACE:
		return BACK_FACE
	case BACK_FACE:
		return FRONT_FACE
	case TOP_FACE:
		return BOTTOM_FACE
	case BOTTOM_FACE:
		return TOP_FACE
	case LEFT_FACE:
		return RIGHT_FACE
	case RIGHT_FACE:
		return LEFT_FACE
	default:
		panic(face)
	}
}

func (face Face) GetNeighbor(direction FacingDirection) Face{
	switch direction {
	case FACING_UP:
		switch face {
		case TOP_FACE:
			return BACK_FACE
		case BACK_FACE:
			return BOTTOM_FACE
		case BOTTOM_FACE:
			return FRONT_FACE
		default:
			return TOP_FACE
		}
	case FACING_DOWN:
		return face.GetNeighbor(FACING_UP).Opposite()
	case FACING_LEFT:
		switch face {
		case RIGHT_FACE:
			return FRONT_FACE
		case LEFT_FACE:
			return BACK_FACE
		default:
			return LEFT_FACE
		}
	case FACING_RIGHT:
		return face.GetNeighbor(FACING_LEFT).Opposite()
	default:
		panic(direction)
	}
}

func (face Face) DirectionOf(other Face) FacingDirection {
	for _, direction := range FacingDirections {
		if face.GetNeighbor(direction) == other {
			return direction
		}
	}
	panic([2]Face{face, other})
}
