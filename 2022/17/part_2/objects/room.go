package objects

type Room struct {
	Contents []uint8
	Width    int
}

func (room *Room) Overlaps(block *Block) bool {
	roomHeight := len(room.Contents)
	for i, blockRow := range block.Contents {
		roomIndex := block.Y + i
		if roomIndex >= roomHeight {
			return false
		}
		roomRow := room.Contents[roomIndex]
		if blockRow&roomRow != 0 {
			return true
		}
	}
	return false
}

func (room *Room) ensureHeight(height int) {
	roomHeight := len(room.Contents)
	if height > roomHeight {
		rows := make([]uint8, height-roomHeight)
		room.Contents = append(room.Contents, rows...)
	}
}

func (room *Room) Place(block *Block) {
	blockYEnd := block.Y + len(block.Contents)
	room.ensureHeight(blockYEnd)
	targetRows := room.Contents[block.Y:blockYEnd]
	for i := range targetRows {
		targetRows[i] |= block.Contents[i]
	}
}

func MakeRoom(width int) Room {
	if width > 8 {
		panic(width)
	}
	return Room{Width: width}
}
