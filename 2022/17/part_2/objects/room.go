package objects

import (
	"math"
)

type Room struct {
	contents []uint8
	Width    int
	YOffset  int
}

func (room *Room) GetHeight() int {
	return room.YOffset + len(room.contents)
}

func (room *Room) Overlaps(block *Block) bool {
	roomHeight := room.GetHeight()
	for i, blockRow := range block.Contents {
		roomIndex := block.Y + i
		if roomIndex < room.YOffset {
			return true
		}
		if roomIndex >= roomHeight {
			return false
		}
		roomRow := room.contents[roomIndex-room.YOffset]
		if blockRow&roomRow != 0 {
			return true
		}
	}
	return false
}

func (room *Room) ensureHeight(height int) {
	roomHeight := room.GetHeight()
	if height > roomHeight {
		rows := make([]uint8, height-roomHeight)
		room.contents = append(room.contents, rows...)
	}
}

func (room *Room) getSlice(start int, end int) []uint8 {
	start -= room.YOffset
	end -= room.YOffset
	return room.contents[start:end]
}

func (room *Room) moveBottom(y int) {
	subContents := room.getSlice(y, room.GetHeight())
	room.YOffset = y
	room.contents = append(room.contents[:0], subContents...)
}

func (room *Room) Place(block *Block) {
	blockYEnd := block.Y + len(block.Contents)
	room.ensureHeight(blockYEnd)
	targetRows := room.getSlice(block.Y, blockYEnd)
	lastFullRow := room.YOffset - 1
	fullRow := uint8(math.Pow(2, float64(room.Width)) - 1)
	for i := range targetRows {
		targetRows[i] |= block.Contents[i]
		if targetRows[i] == fullRow {
			lastFullRow = block.Y + i
		}
	}
	if lastFullRow >= room.YOffset {
		room.moveBottom(lastFullRow + 1)
	}
}

func MakeRoom(width int) Room {
	if width > 8 {
		panic(width)
	}
	return Room{Width: width}
}
