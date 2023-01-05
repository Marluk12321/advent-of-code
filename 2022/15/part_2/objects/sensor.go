package objects

type Sensor struct {
	Position         Position
	closestBeacon    Beacon
	DistanceToBeacon int
}

func (sensor *Sensor) SetBeacon(beacon Beacon) {
	sensor.closestBeacon = beacon
	sensor.DistanceToBeacon = sensor.Position.DistanceTo(beacon.Position)
}

func (sensor *Sensor) GetBeacon(beacon Beacon) Beacon {
	return sensor.closestBeacon
}

func MakeSensor(text string) Sensor {
	positionText := text[len("Sensor at "):]
	return Sensor{
		Position: MakePosition(positionText),
	}
}
