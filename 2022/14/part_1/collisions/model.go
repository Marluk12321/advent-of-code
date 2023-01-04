package collisions

import (
	"2022/14/part_1/world"
)

type CollisionModel struct {
	columns map[int]*CollisionColumn
	Min     world.Position
	Max     world.Position
}

func setBounds(model *CollisionModel, formations []world.RockFormation) {
	model.Min.X = formations[0].Min.X
	model.Min.Y = formations[0].Min.Y
	model.Max.X = formations[0].Max.X
	model.Max.Y = formations[0].Max.Y
	for i := range formations[1:] {
		formation := &formations[i+1]
		if formation.Min.X < model.Min.X {
			model.Min.X = formation.Min.X
		}
		if formation.Max.X > model.Max.X {
			model.Max.X = formation.Max.X
		}
		if formation.Min.Y < model.Min.Y {
			model.Min.Y = formation.Min.Y
		}
		if formation.Max.Y > model.Max.Y {
			model.Max.Y = formation.Max.Y
		}
	}
}

func makeColumns(model *CollisionModel, formations []world.RockFormation) {
	for x := model.Min.X; x <= model.Max.X; x++ {
		column := makeColumn(x, formations)
		if len(column.segments) > 0 {
			model.columns[x] = &column
		}
	}
}

func MakeCollisionModel(formations []world.RockFormation) CollisionModel {
	model := CollisionModel{columns: map[int]*CollisionColumn{}}
	setBounds(&model, formations)
	makeColumns(&model, formations)
	return model
}

func (model *CollisionModel) GetDestination(position world.Position) world.Position {
	destinationY := model.Max.Y + 1
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
		column = &CollisionColumn{x: position.X}
		model.columns[position.X] = column
	}
	column.add(position.Y)
}
