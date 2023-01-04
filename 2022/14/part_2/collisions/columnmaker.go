package collisions

import (
	"2022/14/part_2/world"
	"sort"
)

func makeSegment(point1 int, point2 int) CollisionSegment {
	if point2 < point1 {
		point1, point2 = point2, point1
	}
	return CollisionSegment{
		start: point1,
		end:   point2,
	}
}

func findItersection(vertex1 world.Position, vertex2 world.Position, x int) (*CollisionSegment, bool) {
	xLine := makeSegment(vertex1.X, vertex2.X)
	if x >= xLine.start && x <= xLine.end {
		segment := makeSegment(vertex1.Y, vertex2.Y)
		return &segment, true
	} else {
		return nil, false
	}
}

func makeSegments(formation *world.RockFormation, x int) []CollisionSegment {
	var segments []CollisionSegment
	if x >= formation.Min.X && x <= formation.Max.X {
		for j := 0; j < len(formation.Vertices)-1; j++ {
			vertex1 := formation.Vertices[j]
			vertex2 := formation.Vertices[j+1]
			segment, found := findItersection(vertex1, vertex2, x)
			if found {
				segments = append(segments, *segment)
			}
		}
	}
	return segments
}

func optimize(column *CollisionColumn) {
	if len((*column)) < 2 {
		return
	}
	sort.Slice(*column, func(i, j int) bool {
		segment1 := &(*column)[i]
		segment2 := &(*column)[j]
		return segment1.start < segment2.start ||
			(segment1.start == segment2.start && segment1.end < segment2.end)
	})

	var reduced []CollisionSegment
	workingSegment := &(*column)[0]
	for i := 1; i < len((*column)); i++ {
		segment := &(*column)[i]
		if segment.start > workingSegment.end+1 {
			reduced = append(reduced, *workingSegment)
			workingSegment = segment
		} else if segment.end > workingSegment.end {
			workingSegment.end = segment.end
		}
	}
	reduced = append(reduced, *workingSegment)
	(*column) = reduced
}

func makeColumn(formations []world.RockFormation, x int) CollisionColumn {
	var column CollisionColumn
	for i := range formations {
		formation := &formations[i]
		segments := makeSegments(formation, x)
		column = append(column, segments...)
	}
	optimize(&column)
	return column
}
