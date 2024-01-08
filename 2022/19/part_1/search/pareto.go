package search

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

func removeDominated(states []State, dominator State) []State {
	if len(states) == 0 {
		return states
	}
	filtered := []State{}
	for _, state := range states {
		if !isDominatedBy(state, dominator) {
			filtered = append(filtered, state)
		}
	}
	return filtered
}

func (finder *Finder) saveEvaluated(remaining_minutes int, state State) {
	currentlyBest := finder.bestEvaluated[remaining_minutes]
	filteredBest := removeDominated(currentlyBest, state)
	filteredBest = append(filteredBest, state)
	finder.bestEvaluated[remaining_minutes] = filteredBest
}

func (finder Finder) evaluatedBetter(remaining_minutes int, state State) bool {
	for _, evaluatedState := range finder.bestEvaluated[remaining_minutes] {
		if isDominatedBy(state, evaluatedState) {
			return true
		}
	}
	return false
}
