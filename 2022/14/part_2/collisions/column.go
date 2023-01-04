package collisions

type CollisionSegment struct {
	start int
	end   int
}

type CollisionColumn []CollisionSegment

func (column CollisionColumn) findBlockerIndex(y int) (int, bool) {
	if len(column) == 0 {
		return -1, false
	}

	lastIndex := len(column) - 1
	lastSegment := column[lastIndex]
	if y > lastSegment.end {
		return -1, false
	} else if y >= lastSegment.start {
		return lastIndex, true
	}

	column = column[:lastIndex]
	offset := 0
	bestIndex := lastIndex
	for len(column) > 0 {
		index := len(column) / 2
		segment := column[index]
		if y < segment.start {
			bestIndex = index + offset
			column = column[:index]
		} else if y <= segment.end {
			return index + offset, true
		} else {
			column = column[index+1:]
			offset += index + 1
		}
	}
	return bestIndex, true
}

func (column CollisionColumn) findBlocker(y int) *CollisionSegment {
	index, found := column.findBlockerIndex(y)
	if found {
		return &column[index]
	} else {
		return nil
	}
}

func (column *CollisionColumn) insert(index int, segment CollisionSegment) {
	switch index {
	case len((*column)):
		(*column) = append((*column), segment)
	case 0:
		(*column) = append([]CollisionSegment{segment}, (*column)...)
	default:
		(*column) = append((*column)[:index+1], (*column)[index:]...)
		(*column)[index] = segment
	}
}

func (column *CollisionColumn) add(y int) {
	index, found := column.findBlockerIndex(y)
	if !found {
		(*column) = append((*column), CollisionSegment{
			start: y,
			end:   y,
		})
	} else if (*column)[index].start > y+1 {
		column.insert(index, CollisionSegment{
			start: y,
			end:   y,
		})
	} else {
		(*column)[index].start = y
	}
}
