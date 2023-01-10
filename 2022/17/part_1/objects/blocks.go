package objects

type Block struct {
	Position    XY
	partOffsets []XY
	Size        XY
}

func (block *Block) GetParts() []XY {
	parts := make([]XY, len(block.partOffsets))
	for i, offset := range block.partOffsets {
		parts[i] = block.Position.Plus(offset)
	}
	return parts
}

func calcSize(offsets []XY) XY {
	min := XY{
		X: offsets[0].X,
		Y: offsets[1].Y,
	}
	max := XY{
		X: offsets[0].X,
		Y: offsets[1].Y,
	}
	for _, offset := range offsets {
		if offset.X < min.X {
			min.X = offset.X
		} else if offset.X > max.X {
			max.X = offset.X
		}
		if offset.Y < min.Y {
			min.Y = offset.Y
		} else if offset.Y > max.Y {
			max.Y = offset.Y
		}
	}
	return XY{
		X: max.X - min.X + 1,
		Y: max.Y - min.Y + 1,
	}
}

func makeBlock(position XY, offsets []XY) Block {
	return Block{
		Position:    position,
		partOffsets: offsets,
		Size:        calcSize(offsets),
	}
}

func makeHorizontal(position XY) Block {
	return makeBlock(position, []XY{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	})
}

func makePlus(position XY) Block {
	return makeBlock(position, []XY{
		{1, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{1, 2},
	})
}

func makeCorner(position XY) Block {
	return makeBlock(position, []XY{
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
		{2, 2},
	})
}

func makeVertical(position XY) Block {
	return makeBlock(position, []XY{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	})
}

func makeSquare(position XY) Block {
	return makeBlock(position, []XY{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
	})
}

type BlockType func(XY) Block

var BlockTypes = [5]BlockType{
	makeHorizontal,
	makePlus,
	makeCorner,
	makeVertical,
	makeSquare,
}
