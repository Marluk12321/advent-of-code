package objects

type Sensor struct {
	Position      Position
	ClosestBeacon Beacon
}

func MakeSensor(text string) Sensor {
	positionText := text[len("Sensor at "):]
	return Sensor{
		Position: MakePosition(positionText),
	}
}
