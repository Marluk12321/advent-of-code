package monkeys

type ItemQueue []int

func (queue *ItemQueue) add(value int) {
	(*queue) = append((*queue), value)
}

func (queue *ItemQueue) putBack(value int) {
	(*queue) = append([]int{value}, (*queue)...)
}

func (queue *ItemQueue) first() int {
	return (*queue)[0]
}

func (queue *ItemQueue) take() int {
	value := queue.first()
	(*queue) = (*queue)[1:]
	return value
}

func (queue *ItemQueue) anyLeft() bool {
	return len(*queue) > 0
}
