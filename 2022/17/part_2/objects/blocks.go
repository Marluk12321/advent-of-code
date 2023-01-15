package objects

type Block struct {
	X        int
	Y        int
	Width    int
	Contents []uint8
}

func (block *Block) Shift(amount int) {
	if amount == 0 {
		return
	}
	if amount < 0 {
		amount := uint8(-amount)
		for i := range block.Contents {
			block.Contents[i] >>= amount
		}
	} else {
		amount := uint8(amount)
		for i := range block.Contents {
			block.Contents[i] <<= amount
		}
	}
	block.X += amount
}

func makeBlock(x int, y int) Block {
	block := Block{
		Y: y,
	}
	block.Shift(x)
	return block
}

func makeHorizontal(x int, y int) Block {
	block := makeBlock(x, y)
	block.Contents = []uint8{15}
	block.Width = 4
	return block
}

func makePlus(x int, y int) Block {
	block := makeBlock(x, y)
	block.Contents = []uint8{2, 7, 2}
	block.Width = 3
	return block
}

func makeCorner(x int, y int) Block {
	block := makeBlock(x, y)
	block.Contents = []uint8{7, 4, 4}
	block.Width = 3
	return block
}

func makeVertical(x int, y int) Block {
	block := makeBlock(x, y)
	block.Contents = []uint8{1, 1, 1, 1}
	block.Width = 1
	return block
}

func makeSquare(x int, y int) Block {
	block := makeBlock(x, y)
	block.Contents = []uint8{3, 3}
	block.Width = 2
	return block
}

type BlockType func(int, int) Block

var BlockTypes = [5]BlockType{
	makeHorizontal,
	makePlus,
	makeCorner,
	makeVertical,
	makeSquare,
}
