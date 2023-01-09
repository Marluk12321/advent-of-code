package optimization

import (
	"2022/16/part_2/graph"
	"2022/16/part_2/valve"
	"sort"
)

type State struct {
	currentValves              [2]string
	remainingTimes             [2]int
	closedValves               []string
	pressureReleased           int
	optimisticReleasePotential int
}

func nonZeroValves(valves []valve.Valve) []string {
	var nonZeroValves []string
	for i := range valves {
		valve := &valves[i]
		if valve.FlowRate > 0 {
			nonZeroValves = append(nonZeroValves, valve.Name)
		}
	}
	return nonZeroValves
}

func makeInitialState(valves []valve.Valve, startValve string, availableTime int) State {
	return State{
		currentValves:  [2]string{startValve, startValve},
		remainingTimes: [2]int{availableTime, availableTime},
		closedValves:   nonZeroValves(valves),
	}
}

func openValves(closedValves []string, toOpen [2]string) []string {
	newClosedValves := make([]string, 0, len(closedValves)-len(toOpen))
	shouldOpen := map[string]bool{toOpen[0]: true, toOpen[1]: true}
	for _, name := range closedValves {
		if !shouldOpen[name] {
			newClosedValves = append(newClosedValves, name)
		}
	}
	return newClosedValves
}

func calcReleasePotential(state *State, flowRates map[string]int) {
	rateList := make([]int, 0, len(state.closedValves))
	for _, name := range state.closedValves {
		rateList = append(rateList, flowRates[name])
	}
	sort.Ints(rateList)

	releasePotential := 0
	effectiveTimes := [2]int{
		state.remainingTimes[0],
		state.remainingTimes[1],
	}
	for i := len(rateList) - 1; i >= 0; i-- {
		var moreTimeIndex int
		if effectiveTimes[0] > effectiveTimes[1] {
			moreTimeIndex = 0
		} else {
			moreTimeIndex = 1
		}

		if effectiveTimes[moreTimeIndex] < 2 {
			break
		}
		effectiveTimes[moreTimeIndex] -= 2
		flowRate := rateList[i]
		releasePotential += flowRate * effectiveTimes[moreTimeIndex]
	}
	state.optimisticReleasePotential = releasePotential
}

func makeNextState(currentState *State, nextValves [2]string, graph *graph.Graph) State {
	remainingTimes := [2]int{}
	copy(remainingTimes[:], currentState.remainingTimes[:])
	nextPressureReleased := currentState.pressureReleased
	for i := 0; i < len(nextValves); i++ {
		currentValve, nextValve := currentState.currentValves[i], nextValves[i]
		if nextValve != currentValve {
			remainingTimes[i] -= graph.GetDistance(currentValve, nextValve) + 1
			nextPressureReleased += graph.FlowRates[nextValve] * remainingTimes[i]
		}
	}
	nextClosedValves := openValves(currentState.closedValves, nextValves)

	nextState := State{
		currentValves:    nextValves,
		remainingTimes:   remainingTimes,
		closedValves:     nextClosedValves,
		pressureReleased: nextPressureReleased,
	}
	calcReleasePotential(&nextState, graph.FlowRates)
	return nextState
}
