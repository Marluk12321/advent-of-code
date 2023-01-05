package search

import "sort"

type KnownSegment struct {
	Start int
	End   int
}

func makeSegment(point int, radius int) KnownSegment {
	return KnownSegment{
		Start: point - radius,
		End:   point + radius,
	}
}

func optimize(column *[]KnownSegment) {
	if len((*column)) < 2 {
		return
	}
	sort.Slice(*column, func(i, j int) bool {
		segment1 := &(*column)[i]
		segment2 := &(*column)[j]
		return segment1.Start < segment2.Start ||
			(segment1.Start == segment2.Start && segment1.End < segment2.End)
	})

	var reduced []KnownSegment
	workingSegment := &(*column)[0]
	for i := 1; i < len((*column)); i++ {
		segment := &(*column)[i]
		if segment.Start > workingSegment.End+1 {
			reduced = append(reduced, *workingSegment)
			workingSegment = segment
		} else if segment.End > workingSegment.End {
			workingSegment.End = segment.End
		}
	}
	reduced = append(reduced, *workingSegment)
	(*column) = reduced
}
