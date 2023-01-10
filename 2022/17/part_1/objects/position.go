package objects

type XY struct {
	X int
	Y int
}

func (position XY) Plus(other XY) XY {
	return XY{
		X: position.X + other.X,
		Y: position.Y + other.Y,
	}
}
