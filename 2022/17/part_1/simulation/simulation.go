package simulation

import (
	"2022/17/part_1/objects"
)

func spawn(blockType objects.BlockType, room *objects.Room) objects.Block {
	position := objects.XY{
		X: 2,
		Y: room.GetHeight() + 3,
	}
	return blockType(position)
}

func store(room *objects.Room, block *objects.Block) {
	for _, part := range block.GetParts() {
		room.Place(part)
	}
}

func inCollision(room *objects.Room, block *objects.Block) bool {
	if block.Position.Y >= room.GetHeight() {
		return false
	}
	for _, part := range block.GetParts() {
		if room.IsOccupied(part) {
			return true
		}
	}
	return false
}

func blow(room *objects.Room, block *objects.Block, jetDirection byte) {
	originalBlockX := block.Position.X
	switch jetDirection {
	case '<':
		if block.Position.X == 0 {
			return
		}
		block.Position.X--
	case '>':
		if block.Position.X+block.Size.X == room.Width {
			return
		}
		block.Position.X++
	default:
		panic(jetDirection)
	}
	if inCollision(room, block) {
		block.Position.X = originalBlockX
	}
}

func fall(room *objects.Room, block *objects.Block) {
	if block.Position.Y == 0 {
		return
	}
	block.Position.Y--
	if inCollision(room, block) {
		block.Position.Y++
	}
}

func Simulate(room *objects.Room, blockTypes *[5]objects.BlockType, jetPattern string, spawnLimit int) {
	typeIndex := 0
	jetIndex := 0
	for i := 0; i < spawnLimit; i++ {
		blockType := blockTypes[typeIndex]
		block := spawn(blockType, room)
		for {
			jetDirection := jetPattern[jetIndex]
			blow(room, &block, jetDirection)
			previousBlockY := block.Position.Y
			fall(room, &block)
			jetIndex = (jetIndex + 1) % len(jetPattern)
			if block.Position.Y == previousBlockY {
				break
			}
		}
		store(room, &block)
		typeIndex = (typeIndex + 1) % len(blockTypes)
	}
}
