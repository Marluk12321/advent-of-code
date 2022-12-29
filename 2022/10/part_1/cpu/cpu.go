package cpu

import (
	"strconv"
	"strings"
)

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
			cpu.X += value
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

type CPU struct {
	X                 int
	Cycle             int
	PendingOperations OperationQueue
}

func MakeCPU() CPU {
	return CPU{
		X: 1,
	}
}

func (cpu *CPU) LoadInstruction(instruction string) {
	fields := strings.Fields(instruction)
	var operation Operation
	switch fields[0] {
	case "noop":
		operation = makeNoop()
	case "addx":
		operation = makeAddx(fields[1:])
	default:
		panic(instruction)
	}
	cpu.PendingOperations.push(&operation)
}

func (cpu *CPU) StartCycle() {
	cpu.Cycle++
}

func (cpu *CPU) EndCycle() {
	pendingOperation := cpu.PendingOperations.peek()
	pendingOperation.requiredCycles--
	if pendingOperation.requiredCycles == 0 {
		cpu.PendingOperations.pop().perform(cpu)
	}
}
