package simulation

type State struct {
	jetPatternState   int
	blockFactoryState int
	height            int
}

func (state *State) isEquivalent(other *State) bool {
	return other.jetPatternState == state.jetPatternState &&
		other.blockFactoryState == state.blockFactoryState
}

type Cycle struct {
	startIndex int
	length     int
}

type StartingPoint struct {
	index    int
	distance int
}

func findCycleStartingPoints(states []State) []StartingPoint {
	latestIndex := len(states) - 1
	latestState := &states[latestIndex]
	var startingPoints []StartingPoint
	for i := range states[:latestIndex] {
		state := &states[i]
		if state.isEquivalent(latestState) {
			startingPoints = append(startingPoints, StartingPoint{
				index:    i,
				distance: latestIndex - i,
			})
		}
	}
	return startingPoints
}

func isCycle(states []State, startingPoint StartingPoint) bool {
	latestIndex := len(states) - 1
	for offset := 1; offset < startingPoint.distance; offset++ {
		if offset > startingPoint.index {
			return false
		}
		state1 := &states[startingPoint.index-offset]
		state2 := &states[latestIndex-offset]
		if !state1.isEquivalent(state2) {
			return false
		}
	}
	return true
}

func detectCycles(states []State) []Cycle {
	var detectedCycles []Cycle
	startingPoints := findCycleStartingPoints(states)
	for _, startingPoint := range startingPoints {
		if isCycle(states, startingPoint) {
			detectedCycles = append(detectedCycles, Cycle{
				startIndex: startingPoint.index - startingPoint.distance + 1,
				length:     startingPoint.distance,
			})
		}
	}
	return detectedCycles
}

func getLongest(cycles []Cycle) Cycle {
	longest := cycles[0]
	for _, cycle := range cycles[1:] {
		if cycle.length > longest.length {
			longest = cycle
		}
	}
	return longest
}

func findCycle(states []State) (Cycle, bool) {
	cycles := detectCycles(states)
	if len(cycles) == 0 {
		return Cycle{}, false
	}
	return getLongest(cycles), true
}
