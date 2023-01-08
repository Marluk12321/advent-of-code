package optimization

import (
	"2022/16/part_1/graph"
	"2022/16/part_1/valve"
	"container/heap"
)

func MaxPressureReleased(valves []valve.Valve, startValve string, availableTime int) int {
	graph := graph.MakeGraph(valves)
	initialState := makeInitialState(valves, startValve, availableTime)
	queue := PriorityQueue{}
	heap.Push(&queue, &initialState)

	for queue.Len() > 0 {
		obj := heap.Pop(&queue)
		state := obj.(*State)
		if state.remainingTime == 0 || len(state.closedValves) == 0 {
			return state.pressureReleased
		}
		for name := range state.closedValves {
			nextState := makeNextState(state, name, &graph)
			heap.Push(&queue, &nextState)
		}
	}

	return 0
}
