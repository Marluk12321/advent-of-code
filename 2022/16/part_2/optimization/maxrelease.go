package optimization

import (
	"2022/16/part_2/graph"
	"2022/16/part_2/valve"
	"container/heap"
)

func getOptions(state *State, index int, graph *graph.Graph) []string {
	var options []string
	current := state.currentValves[index]
	remainingTime := state.remainingTimes[index]
	for _, name := range state.closedValves {
		if graph.GetDistance(current, name) <= remainingTime {
			options = append(options, name)
		}
	}
	return options
}

func getNext(state *State, graph *graph.Graph) []State {
	options := [2][]string{
		getOptions(state, 0, graph),
		getOptions(state, 1, graph),
	}
	var nextStates []State
	if len(options[0]) == 0 {
		if len(options[1]) == 0 {
			return nextStates
		} else {
			options[0] = append(options[0], state.currentValves[0])
		}
	}
	if len(options[1]) == 0 {
		options[1] = append(options[1], state.currentValves[1])
	}

	for _, name1 := range options[0] {
		for _, name2 := range options[1] {
			if name1 != name2 {
				nextValves := [2]string{name1, name2}
				nextState := makeNextState(state, nextValves, graph)
				nextStates = append(nextStates, nextState)
			}
		}
	}
	return nextStates
}

func MaxPressureReleased(valves []valve.Valve, startValve string, availableTime int) int {
	graph := graph.MakeGraph(valves)
	initialState := makeInitialState(valves, startValve, availableTime)
	queue := PriorityQueue{}
	heap.Push(&queue, &initialState)

	for queue.Len() > 0 {
		obj := heap.Pop(&queue)
		state := obj.(*State)
		if len(state.closedValves) == 0 {
			return state.pressureReleased
		}
		nextStates := getNext(state, &graph)
		if len(nextStates) == 0 {
			return state.pressureReleased
		}
		for i := range nextStates {
			heap.Push(&queue, &nextStates[i])
		}
	}

	return 0
}
