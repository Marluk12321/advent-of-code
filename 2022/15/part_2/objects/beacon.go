package objects

type Beacon struct {
	Position Position
}

func MakeBeacon(text string) Beacon {
	positionText := text[len("closest beacon is at "):]
	return Beacon{
		Position: MakePosition(positionText),
	}
}
