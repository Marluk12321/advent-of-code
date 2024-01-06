package simulation

import "2022/19/part_1/objects"

func getBuildableRobots(
	constructionCosts objects.ConstructionCosts,
	availableResources objects.Resources,
) []objects.Robot {
	options := []objects.Robot{}
	for robot, cost := range constructionCosts {
		allAvailable := true
		for resource, amount := range cost {
			if availableResources[resource] < amount {
				allAvailable = false
				break
			}
		}
		if allAvailable {
			options = append(options, robot)
		}
	}
	return options
}

func FindBestScore(
	constructionCosts objects.ConstructionCosts,
	state State,
	remaining_minutes int,
) int {
	if remaining_minutes == 0 {
		return state.Resources[objects.OPEN_GEODE]
	}
	nextState := state.getNext()
	remaining_minutes--
	bestScore := FindBestScore(constructionCosts, nextState, remaining_minutes)
	for _, robot := range getBuildableRobots(constructionCosts, state.Resources) {
		nextState = state.startBuilding(robot, constructionCosts[robot]).getNext()
		score := FindBestScore(constructionCosts, nextState, remaining_minutes)
		if score > bestScore {
			bestScore = score
		}
	}
	return bestScore
}
