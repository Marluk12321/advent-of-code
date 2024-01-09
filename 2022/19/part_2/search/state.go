package search

import "2022/19/part_1/objects"

type Robots map[objects.Robot]int

type State struct {
	Robots    Robots
	Resources objects.Resources
}

func copyResources(resources objects.Resources) objects.Resources {
	result := make(objects.Resources, len(resources))
	for resource, amount := range resources {
		result[resource] = amount
	}
	return result
}

func increaseResources(resources objects.Resources, robots Robots) {
	for robot, amount := range robots {
		resources[robot.Resource] += amount
	}
}

func decreaseResources(resources, cost objects.Resources) {
	for resource, amount := range cost {
		resources[resource] -= amount
	}
}

func copyRobots(robots Robots) Robots {
	result := make(Robots, len(robots))
	for robot, amount := range robots {
		result[robot] = amount
	}
	return result
}

func (state State) produce() State {
	nextResources := copyResources(state.Resources)
	increaseResources(nextResources, state.Robots)
	return State{
		Robots:    state.Robots,
		Resources: nextResources,
	}
}

func (state State) buildAndProduce(robot objects.Robot, cost objects.Resources) State {
	nextResources := copyResources(state.Resources)
	decreaseResources(nextResources, cost)
	increaseResources(nextResources, state.Robots)
	nextRobots := copyRobots(state.Robots)
	nextRobots[robot]++
	return State{
		Robots:    nextRobots,
		Resources: nextResources,
	}
}
