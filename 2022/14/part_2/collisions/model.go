package collisions

import (
	"2022/14/part_2/world"
)

type CollisionModel struct {
	columns map[int]*CollisionColumn
	floorY  int
}

func getXBounds(formations []world.RockFormation) (int, int) {
	minX := formations[0].Min.X
	maxX := formations[0].Max.X
	for i := range formations[1:] {
		formation := &formations[i+1]
		if formation.Min.X < minX {
			minX = formation.Min.X
		}
		if formation.Max.X > maxX {
			maxX = formation.Max.X
		}
	}
	return minX, maxX
}

func makeColumns(model *CollisionModel, formations []world.RockFormation) {
	minX, maxX := getXBounds(formations)
	for x := minX; x <= maxX; x++ {
		column := makeColumn(formations, x)
		if len(column) > 0 {
			model.columns[x] = &column
		}
	}
}

func MakeCollisionModel(world *world.World) CollisionModel {
	model := CollisionModel{
		columns: map[int]*CollisionColumn{},
		floorY:  world.FloorY,
	}
	makeColumns(&model, world.RockFormations)
	return model
}

func (model *CollisionModel) GetDestination(position world.Position) world.Position {
	destinationY := model.floorY - 1
	column, exists := model.columns[position.X]
	if exists {
		segment := column.findBlocker(position.Y)
		if segment != nil {
			destinationY = segment.start - 1
		}
	}
	return world.Position{
		X: position.X,
		Y: destinationY,
	}
}

func (model *CollisionModel) Add(position world.Position) {
	column, exists := model.columns[position.X]
	if !exists {
		column = &CollisionColumn{}
		model.columns[position.X] = column
	}
	column.add(position.Y)
}
