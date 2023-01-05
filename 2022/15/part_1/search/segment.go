package search

import "sort"

type EmptySegment struct {
	start int
	end   int
}

func makeSegment(point int, radius int) EmptySegment {
	return EmptySegment{
		start: point - radius,
		end:   point + radius,
	}
}

func (segment *EmptySegment) remove(point int) []EmptySegment {
	if point < segment.start || point > segment.end {
		panic([3]int{segment.start, segment.end, point})
	}
	switch point {
	case segment.start:
		return []EmptySegment{{
			start: segment.start + 1,
			end:   segment.end,
		}}
	case segment.end:
		return []EmptySegment{{
			start: segment.start + 1,
			end:   segment.end,
		}}
	default:
		return []EmptySegment{{
			start: segment.start,
			end:   point - 1,
		}, {
			start: point + 1,
			end:   segment.end,
		}}
	}
}

func optimize(column *[]EmptySegment) {
	if len((*column)) < 2 {
		return
	}
	sort.Slice(*column, func(i, j int) bool {
		segment1 := &(*column)[i]
		segment2 := &(*column)[j]
		return segment1.start < segment2.start ||
			(segment1.start == segment2.start && segment1.end < segment2.end)
	})

	var reduced []EmptySegment
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
