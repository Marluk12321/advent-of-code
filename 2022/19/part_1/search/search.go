package search

import "2022/19/part_1/objects"

type Finder struct {
	constructionCosts objects.ConstructionCosts
	maxCosts objects.Resources
	scoringResource objects.Resource
	bestEvaluated map[int][]State
}

func makeFinder(
	constructionCosts objects.ConstructionCosts,
	scoringResource objects.Resource,
) Finder {
	maxCosts := objects.Resources{}
	for _, cost := range constructionCosts {
		for resource, amount := range cost {
			if maxCosts[resource] < amount {
				maxCosts[resource] = amount
			}
		}
	}
	return Finder{
		constructionCosts: constructionCosts,
		maxCosts: maxCosts,
		scoringResource: scoringResource,
		bestEvaluated: map[int][]State{},
	}
}

func canAfford(cost objects.Resources, availableResources objects.Resources) bool {
	for resource, amount := range cost {
		if availableResources[resource] < amount {
			return false
		}
	}
	return true
}

func (finder Finder) getBuildOptions(state State, remaining_minutes int) []objects.Robot {
	options := []objects.Robot{}
	if remaining_minutes <= 1 {
		// any built robot doesn't have time to produce
		return options
	}

	scoreProducer := objects.RobotForResource[finder.scoringResource]
	if canAfford(finder.constructionCosts[scoreProducer], state.Resources) {
		options = append(options, scoreProducer)
	}
	if remaining_minutes <= 3 {
		// only generating score resource is useful in the last 2 minutes
		return options
	}

	if remaining_minutes <= 4 && len(options) == 1 {
		// builing any other robots could only be useful if score producer can't be built yet
		return options
	}

	if remaining_minutes <= 5 {
		// it's only useful to generate score resource and score producer's requirements
		for resource, amount := range finder.constructionCosts[scoreProducer] {
			resourceProducer := objects.RobotForResource[resource]
			resourceProducerCost := finder.constructionCosts[resourceProducer]
			if amount > state.Robots[resourceProducer] &&
					canAfford(resourceProducerCost, state.Resources) {
				options = append(options, resourceProducer)
			}
		}
		return options
	}

	for robot, cost := range finder.constructionCosts {
		if robot == scoreProducer {
			continue
		}
		resource := robot.Resource
		if finder.maxCosts[resource] > state.Robots[robot] && canAfford(cost, state.Resources) {
			options = append(options, robot)
		}
	}
	return options
}

func (finder Finder) findBestScore(remaining_minutes int, state State) int {
	if remaining_minutes == 0 {
		return state.Resources[finder.scoringResource]
	}
	if finder.evaluatedBetter(remaining_minutes, state) {
		return 0
	}
	bestScore := 0
	for _, robot := range finder.getBuildOptions(state, remaining_minutes) {
		nextState := state.buildAndProduce(robot, finder.constructionCosts[robot])
		score := finder.findBestScore(remaining_minutes - 1, nextState)
		if score > bestScore {
			bestScore = score
		}
	}
	nextState := state.produce()
	score := finder.findBestScore(remaining_minutes - 1, nextState)
	if score > bestScore {
		bestScore = score
	}
	finder.saveEvaluated(remaining_minutes, state)
	return bestScore
}

func FindBestScore(
	constructionCosts objects.ConstructionCosts,
	scoringResource objects.Resource,
	minute_limit int,
	initialState State,
) int {
	finder := makeFinder(constructionCosts, scoringResource)
	return finder.findBestScore(minute_limit, initialState)
}
