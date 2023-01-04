package sand

import (
	"2022/14/part_2/collisions"
	"2022/14/part_2/world"
)

func spawn(collisionModel *collisions.CollisionModel, spawnPosition world.Position) world.Position {
	destination := collisionModel.GetDestination(spawnPosition)
	if destination.Y < spawnPosition.Y {
		return destination
	}

	leftSpawnPosition := world.Position{
		X: destination.X - 1,
		Y: destination.Y + 1,
	}
	leftDestination := spawn(collisionModel, leftSpawnPosition)
	if leftDestination.Y >= leftSpawnPosition.Y {
		return leftDestination
	}

	rightSpawnPosition := world.Position{
		X: destination.X + 1,
		Y: destination.Y + 1,
	}
	rightDestination := spawn(collisionModel, rightSpawnPosition)
	if rightDestination.Y >= rightSpawnPosition.Y {
		return rightDestination
	}

	return destination
}

func Fill(world *world.World, spawnPosition world.Position) {
	collisionModel := collisions.MakeCollisionModel(world)
	for {
		sandPosition := spawn(&collisionModel, spawnPosition)
		world.SandPositions = append(world.SandPositions, sandPosition)
		collisionModel.Add(sandPosition)
		if sandPosition.Y == spawnPosition.Y {
			break
		}
	}
}
