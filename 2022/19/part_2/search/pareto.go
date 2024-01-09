package search

import "2022/19/part_1/objects"

func (finder Finder) theoreticalWorstScore(remaining_minutes int, state State) int {
	score := state.Resources[finder.scoringResource]
	scoreProducer := objects.RobotForResource[finder.scoringResource]
	score += remaining_minutes * state.Robots[scoreProducer]
	return score
}

var additionalScoreCache = map[int]int{1: 0}

func calculateAddidionalScore(remaining_minutes int) int {
	additionalScore, found := additionalScoreCache[remaining_minutes]
	if !found {
		reduced_minutes := remaining_minutes - 1
		additionalScore = reduced_minutes + calculateAddidionalScore(reduced_minutes)
		additionalScoreCache[remaining_minutes] = additionalScore
	}
	return additionalScore
}

func (finder Finder) theoreticalBestScore(remaining_minutes int, state State) int {
	score := finder.theoreticalWorstScore(remaining_minutes, state)
	return score + calculateAddidionalScore(remaining_minutes)
}

func isDominatedBy(state State, dominator State) bool {
	if len(state.Robots) > len(dominator.Robots) ||
		len(state.Resources) > len(dominator.Resources) {
		return false
	}
	isDominated := false
	for robot, amount := range state.Robots {
		dominatorAmount := dominator.Robots[robot]
		if amount > dominatorAmount {
			return false
		} else if amount < dominatorAmount {
			isDominated = true
		}
	}
	for resource, amount := range state.Resources {
		dominatorAmount := dominator.Resources[resource]
		if amount > dominatorAmount {
			return false
		} else if amount < dominatorAmount {
			isDominated = true
		}
	}
	// has more of the shared robot/resource types
	return isDominated ||
		// has other robot types besides equal shared ones
		len(dominator.Robots) > len(state.Robots) ||
		// has other resource types besides equal shared ones
		len(dominator.Resources) > len(state.Resources)
}

func (finder Finder) withoutDominated(remaining_minutes int, dominator State) []State {
	currentlyBest := finder.dominatingStates[remaining_minutes]
	if len(currentlyBest) == 0 {
		return currentlyBest
	}
	filtered := []State{}
	for _, state := range currentlyBest {
		if !isDominatedBy(state, dominator) {
			filtered = append(filtered, state)
		}
	}
	return filtered
}

func (finder *Finder) updateDominatingStates(remaining_minutes int, state State) {
	filteredBest := finder.withoutDominated(remaining_minutes, state)
	filteredBest = append(filteredBest, state)
	finder.dominatingStates[remaining_minutes] = filteredBest
}

func (finder Finder) seenDominatingState(remaining_minutes int, state State) bool {
	for _, evaluatedState := range finder.dominatingStates[remaining_minutes] {
		if isDominatedBy(state, evaluatedState) {
			return true
		}
	}
	return false
}
