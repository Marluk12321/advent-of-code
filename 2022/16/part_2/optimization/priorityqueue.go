package optimization

type PriorityQueue []*State

func (queue *PriorityQueue) Len() int {
	return len(*queue)
}

func (queue *PriorityQueue) Less(i int, j int) bool {
	stateI := (*queue)[i]
	stateJ := (*queue)[j]
	scoreI := stateI.pressureReleased + stateI.optimisticReleasePotential
	scoreJ := stateJ.pressureReleased + stateJ.optimisticReleasePotential
	return scoreI > scoreJ
}

func (queue *PriorityQueue) Pop() any {
	index := queue.Len() - 1
	state := &(*queue)[index]
	(*queue) = (*queue)[:index]
	return *state
}

func (queue *PriorityQueue) Push(obj any) {
	state := obj.(*State)
	(*queue) = append((*queue), state)
}

func (queue *PriorityQueue) Swap(i int, j int) {
	(*queue)[i], (*queue)[j] = (*queue)[j], (*queue)[i]
}
