package system

import "strconv"

type Operation struct {
	requiredCycles int
	perform        func(*CPU)
}

func makeNoop() Operation {
	return Operation{
		requiredCycles: 1,
		perform:        func(cpu *CPU) {},
	}
}

func makeAddx(additionalFields []string) Operation {
	value, err := strconv.Atoi(additionalFields[0])
	if err != nil {
		panic(err)
	}
	return Operation{
		requiredCycles: 2,
		perform: func(cpu *CPU) {
			cpu.x += value
		},
	}
}

type OperationQueue []Operation

func (queue *OperationQueue) push(operation *Operation) {
	(*queue) = append((*queue), *operation)
}

func (queue *OperationQueue) peek() *Operation {
	return &(*queue)[0]
}

func (queue *OperationQueue) pop() *Operation {
	operation := queue.peek()
	(*queue) = (*queue)[1:]
	return operation
}
