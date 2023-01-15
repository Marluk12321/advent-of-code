package simulation

import (
	"2022/17/part_2/objects"
)

func spawn(blockType objects.BlockType, room *objects.Room, jetPattern *JetPattern) objects.Block {
	block := blockType(0, len(room.Contents))
	xOffset := 2
	minXOffest := 0
	maxXOffset := room.Width - block.Width
	for i := 0; i < 3; i++ {
		xOffset += jetPattern.NextXOffset()
		if xOffset < minXOffest {
			xOffset = minXOffest
		} else if xOffset > maxXOffset {
			xOffset = maxXOffset
		}
	}
	block.Shift(xOffset)
	return block
}

func blow(room *objects.Room, block *objects.Block, jetPattern *JetPattern) {
	xOffset := jetPattern.NextXOffset()
	if xOffset < 0 {
		if block.X == 0 {
			return
		}
	} else if block.X+block.Width >= room.Width {
		return
	}
	block.Shift(xOffset)
	if room.Overlaps(block) {
		block.Shift(-xOffset)
	}
}

func fall(room *objects.Room, block *objects.Block) {
	if block.Y == 0 {
		return
	}
	block.Y--
	if room.Overlaps(block) {
		block.Y++
	}
}

func extrapolateFinalHeight(states []State, cycle *Cycle, spawnLimit int) int {
	heightPreCycle := states[cycle.startIndex].height
	heightPostCycle := states[cycle.startIndex+cycle.length].height
	heightPerCycle := heightPostCycle - heightPreCycle

	spawnsSinceCycle := spawnLimit - cycle.startIndex
	fullCycles := spawnsSinceCycle / cycle.length

	spawnsRemaining := spawnsSinceCycle % cycle.length
	heightRemaining := states[cycle.startIndex+spawnsRemaining].height - heightPreCycle

	return heightPreCycle + fullCycles*heightPerCycle + heightRemaining
}

func CalcHeight(room *objects.Room, blockTypes *[5]objects.BlockType, jetPattern *JetPattern, spawnLimit int) int {
	typeIndex := 0
	var states []State
	for i := 0; i < spawnLimit; i++ {
		states = append(states, State{
			jetPatternState:   jetPattern.index,
			blockFactoryState: typeIndex,
			height:            len(room.Contents),
		})
		if cycle, found := findCycle(states); found {
			return extrapolateFinalHeight(states, &cycle, spawnLimit)
		}

		blockType := blockTypes[typeIndex]
		block := spawn(blockType, room, jetPattern)
		for {
			blow(room, &block, jetPattern)
			previousBlockHeight := block.Y
			fall(room, &block)
			if block.Y == previousBlockHeight {
				break
			}
		}
		room.Place(&block)
		typeIndex = (typeIndex + 1) % len(blockTypes)
	}
	return len(room.Contents)
}
