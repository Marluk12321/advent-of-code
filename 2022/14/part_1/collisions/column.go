package collisions

import (
	"2022/14/part_1/world"
	"sort"
)

type CollisionSegment struct {
	start int
	end   int
}

type CollisionColumn struct {
	x        int
	segments []CollisionSegment
}

func (column *CollisionColumn) loadVertical(x int, y1 int, y2 int) {
	if x != column.x {
		return
	}
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	segment := CollisionSegment{
		start: y1,
		end:   y2,
	}
	column.segments = append(column.segments, segment)
}

func (column *CollisionColumn) loadHorizontal(y int, x1 int, x2 int) {
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	if column.x >= x1 && column.x <= x2 {
		segment := CollisionSegment{
			start: y,
			end:   y,
		}
		column.segments = append(column.segments, segment)
	}
}

func (column *CollisionColumn) loadLine(vertex1 *world.Position, vertex2 *world.Position) {
	if vertex1.X == vertex2.X {
		column.loadVertical(vertex1.X, vertex1.Y, vertex2.Y)
	} else if vertex1.Y == vertex2.Y {
		column.loadHorizontal(vertex1.Y, vertex1.X, vertex2.X)
	} else {
		panic([]world.Position{*vertex1, *vertex2})
	}
}

func (column *CollisionColumn) loadFormation(formation *world.RockFormation) {
	if column.x >= formation.Min.X && column.x <= formation.Max.X {
		for j := 0; j < len(formation.Vertices)-1; j++ {
			vertex1 := &formation.Vertices[j]
			vertex2 := &formation.Vertices[j+1]
			column.loadLine(vertex1, vertex2)
		}
	}
}

func (column *CollisionColumn) optimize() {
	if len(column.segments) < 2 {
		return
	}
	sort.Slice(column.segments, func(i, j int) bool {
		segment1 := &column.segments[i]
		segment2 := &column.segments[j]
		return segment1.start < segment2.start ||
			(segment1.start == segment2.start && segment1.end < segment2.end)
	})

	var reduced []CollisionSegment
	workingSegment := &column.segments[0]
	for i := 1; i < len(column.segments); i++ {
		segment := &column.segments[i]
		if segment.start > workingSegment.end+1 {
			reduced = append(reduced, *workingSegment)
			workingSegment = segment
		} else if segment.end > workingSegment.end {
			workingSegment.end = segment.end
		}
	}
	reduced = append(reduced, *workingSegment)
	column.segments = reduced
}

func makeColumn(x int, formations []world.RockFormation) CollisionColumn {
	column := CollisionColumn{x: x}
	for i := range formations {
		formation := &formations[i]
		column.loadFormation(formation)
	}
	column.optimize()
	return column
}

func (column *CollisionColumn) findBlockerIndex(y int) (int, bool) {
	segments := column.segments
	if len(segments) == 0 {
		return -1, false
	}

	lastIndex := len(segments) - 1
	lastSegment := segments[lastIndex]
	if y > lastSegment.end {
		return -1, false
	} else if y >= lastSegment.start {
		return lastIndex, true
	}

	segments = segments[:lastIndex]
	offset := 0
	bestIndex := lastIndex
	for len(segments) > 0 {
		index := len(segments) / 2
		segment := segments[index]
		if y < segment.start {
			bestIndex = index + offset
			segments = segments[:index]
		} else if y <= segment.end {
			return index + offset, true
		} else {
			segments = segments[index+1:]
			offset += index + 1
		}
	}
	return bestIndex, true
}

func (column *CollisionColumn) findBlocker(y int) *CollisionSegment {
	index, found := column.findBlockerIndex(y)
	if found {
		return &column.segments[index]
	} else {
		return nil
	}
}

func (column *CollisionColumn) insert(index int, segment CollisionSegment) {
	switch index {
	case 0:
		column.segments = append([]CollisionSegment{segment}, column.segments...)
	case len(column.segments):
		column.segments = append(column.segments, segment)
	default:
		column.segments = append(column.segments[:index+1], column.segments[index:]...)
		column.segments[index] = segment
	}
}

func (column *CollisionColumn) add(y int) {
	index, found := column.findBlockerIndex(y)
	if !found {
		panic(y)
	}
	if column.segments[index].start > y+1 {
		column.insert(index, CollisionSegment{
			start: y,
			end:   y,
		})
	} else {
		segment := &column.segments[index]
		segment.start = y
	}
}
