package search

import (
	"2022/15/part_2/objects"
)

func distanceToLine(position objects.Position, y int) int {
	if position.Y > y {
		return position.Y - y
	} else {
		return y - position.Y
	}
}

func FindKnownSegments(sensors []objects.Sensor, y int) []KnownSegment {
	var emptySegments []KnownSegment
	for i := range sensors {
		sensor := &sensors[i]
		distanceToLine := distanceToLine(sensor.Position, y)
		distanceToBeacon := sensor.DistanceToBeacon
		if distanceToLine < distanceToBeacon {
			difference := distanceToBeacon - distanceToLine
			segment := makeSegment(sensor.Position.X, difference)
			emptySegments = append(emptySegments, segment)
		}
	}
	optimize(&emptySegments)
	return emptySegments
}
