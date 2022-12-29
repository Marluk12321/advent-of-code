package system

import "strings"

type CPU struct {
	x                 int
	PendingOperations OperationQueue
}

func MakeCPU() CPU {
	return CPU{
		x: 1,
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

func (cpu *CPU) endCycle() {
	pendingOperation := cpu.PendingOperations.peek()
	pendingOperation.requiredCycles--
	if pendingOperation.requiredCycles == 0 {
		cpu.PendingOperations.pop().perform(cpu)
	}
}
