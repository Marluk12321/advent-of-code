package sand

import (
	"2022/14/part_1/collisions"
	"2022/14/part_1/world"
)

func spawn(collisionModel *collisions.CollisionModel, spawnPosition world.Position) world.Position {
	destination := collisionModel.GetDestination(spawnPosition)
	if destination.Y < spawnPosition.Y || destination.Y > collisionModel.Max.Y {
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
	collisionModel := collisions.MakeCollisionModel(world.RockFormations)
	for {
		sandPosition := spawn(&collisionModel, spawnPosition)
		if sandPosition.Y <= spawnPosition.Y {
			panic(sandPosition)
		}
		if sandPosition.Y <= collisionModel.Max.Y {
			world.SandPositions = append(world.SandPositions, sandPosition)
			collisionModel.Add(sandPosition)
		} else {
			break
		}
	}
}
