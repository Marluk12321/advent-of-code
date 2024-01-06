package simulation

import "2022/19/part_1/objects"

type Robots map[objects.Robot]int

type State struct {
	Robots    Robots
	Resources objects.Resources
	Factory   objects.Factory
}

func (state State) duplicate() State {
	result := State{
		Robots:    make(Robots, len(state.Robots)),
		Resources: make(objects.Resources, len(state.Resources)),
		Factory: objects.Factory{
			IsProducing: state.Factory.IsProducing,
			Robot:       state.Factory.Robot,
		},
	}
	for robot, amount := range state.Robots {
		result.Robots[robot] = amount
	}
	for resource, amount := range state.Resources {
		result.Resources[resource] = amount
	}
	return result
}

func (state State) startBuilding(robot objects.Robot, cost objects.Resources) State {
	result := state.duplicate()
	for resource, amount := range cost {
		if currentAmount := result.Resources[resource]; currentAmount < amount {
			panic([2]int{currentAmount, amount})
		}
		result.Resources[resource] -= amount
	}
	result.Factory.IsProducing = true
	result.Factory.Robot = robot
	return result
}

func (state State) getNext() State {
	result := state.duplicate()
	for robot, amount := range state.Robots {
		result.Resources[robot.Resource] += amount
	}
	if state.Factory.IsProducing {
		result.Robots[state.Factory.Robot]++
		result.Factory.IsProducing = false
	}
	return result
}
