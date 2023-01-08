package optimization

import (
	"2022/16/part_1/graph"
	"2022/16/part_1/valve"
	"sort"
)

type State struct {
	currentValve               string
	remainingTime              int
	closedValves               map[string]bool
	pressureReleased           int
	optimisticReleasePotential int
}

func nonZeroValves(valves []valve.Valve) map[string]bool {
	nonZeroValves := map[string]bool{}
	for i := range valves {
		valve := &valves[i]
		if valve.FlowRate > 0 {
			nonZeroValves[valve.Name] = true
		}
	}
	return nonZeroValves
}

func makeInitialState(valves []valve.Valve, startValve string, availableTime int) State {
	return State{
		currentValve:  startValve,
		remainingTime: availableTime,
		closedValves:  nonZeroValves(valves),
	}
}

func openValve(closedValves map[string]bool, toOpen string) map[string]bool {
	newClosedValves := make(map[string]bool, len(closedValves)-1)
	for name := range closedValves {
		if name != toOpen {
			newClosedValves[name] = true
		}
	}
	return newClosedValves
}

func calcReleasePotential(state *State, flowRates map[string]int) {
	rateList := make([]int, len(state.closedValves))
	for name := range state.closedValves {
		rateList = append(rateList, flowRates[name])
	}
	sort.Ints(rateList)

	releasePotential := 0
	effectiveTime := state.remainingTime - 2
	for i := len(rateList) - 1; i >= 0; i-- {
		if effectiveTime < 2 {
			break
		}
		flowRate := rateList[i]
		releasePotential += flowRate * effectiveTime
		effectiveTime -= 2
	}
	state.optimisticReleasePotential = releasePotential
}

func makeNextState(currentState *State, nextValve string, graph *graph.Graph) State {
	moveCost := graph.GetDistance(currentState.currentValve, nextValve) + 1
	effectiveTime := currentState.remainingTime - moveCost
	moveGain := graph.FlowRates[nextValve] * effectiveTime
	nextClosedValves := openValve(currentState.closedValves, nextValve)
	nextState := State{
		currentValve:     nextValve,
		remainingTime:    effectiveTime,
		closedValves:     nextClosedValves,
		pressureReleased: currentState.pressureReleased + moveGain,
	}
	calcReleasePotential(&nextState, graph.FlowRates)
	return nextState
}
