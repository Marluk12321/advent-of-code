package search

import (
	"2022/15/part_1/objects"
)

func distanceToLine(position objects.Position, y int) int {
	closestPosition := objects.Position{
		X: position.X,
		Y: y,
	}
	return position.DistanceTo(closestPosition)
}

func findEmptySegments(sensors []objects.Sensor, y int) []EmptySegment {
	var emptySegments []EmptySegment
	for i := range sensors {
		sensor := &sensors[i]
		distanceToLine := distanceToLine(sensor.Position, y)
		distanceToBeacon := sensor.Position.DistanceTo(sensor.ClosestBeacon.Position)
		if distanceToLine < distanceToBeacon {
			difference := distanceToBeacon - distanceToLine
			segment := makeSegment(sensor.Position.X, difference)
			if sensor.ClosestBeacon.Position.Y == y {
				toRemove := sensor.ClosestBeacon.Position.X
				segments := segment.remove(toRemove)
				emptySegments = append(emptySegments, segments...)
			} else {
				emptySegments = append(emptySegments, segment)
			}
		}
	}
	optimize(&emptySegments)
	return emptySegments
}

func KnownEmptySpaces(sensors []objects.Sensor, y int) int {
	emptySegments := findEmptySegments(sensors, y)
	var emptySpaces int
	for _, segment := range emptySegments {
		emptySpaces += segment.end - segment.start + 1
	}
	return emptySpaces
}
