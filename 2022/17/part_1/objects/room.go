package objects

type Room struct {
	contents []uint8
	Width    int
}

func (room *Room) GetHeight() int {
	return len(room.contents)
}

func (room *Room) checkX(position XY) {
	if position.X >= room.Width {
		panic(position)
	}
}

func (room *Room) ensureHeight(height int) {
	roomHeight := room.GetHeight()
	if height > roomHeight {
		rows := make([]uint8, height-roomHeight)
		room.contents = append(room.contents, rows...)
	}
}

func (room *Room) Place(objects ...XY) {
	for _, object := range objects {
		room.checkX(object)
		room.ensureHeight(object.Y + 1)
		room.contents[object.Y] |= 1 << object.X
	}
}

func (room *Room) IsOccupied(position XY) bool {
	room.checkX(position)
	if position.Y >= room.GetHeight() {
		return false
	}
	row := room.contents[position.Y]
	mask := uint8(1 << position.X)
	return row&mask != 0
}

func MakeRoom(width int) Room {
	if width > 8 {
		panic(width)
	}
	return Room{Width: width}
}
