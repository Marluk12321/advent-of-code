package instructions

import (
	"2022/22/part_2/mapping"
	"strconv"
)

type Instruction interface {
	Apply(cubeMap mapping.Cube, state mapping.State) mapping.State
	String() string
}

type Instructions []Instruction

func flushBuffer(instructions *Instructions, numberBuffer *[]rune) {
	numberText := string(*numberBuffer)
	number, err := strconv.Atoi(numberText)
	if err != nil {
		panic(err)
	}
	instruction := MoveInstruction{
		moveBy: number,
	}
	*instructions = append(*instructions, instruction)
	*numberBuffer = (*numberBuffer)[:0]
}

func BuildInstructions(text string) Instructions {
	instructions := Instructions{}
	numberBuffer := []rune{}
	for _, c := range text {
		switch c {
		case 'L':
			if len(numberBuffer) > 0 {
				flushBuffer(&instructions, &numberBuffer)
			}
			instruction := TurnInstruction{
				turnDirection: LEFT,
			}
			instructions = append(instructions, instruction)
		case 'R':
			if len(numberBuffer) > 0 {
				flushBuffer(&instructions, &numberBuffer)
			}
			instruction := TurnInstruction{
				turnDirection: RIGHT,
			}
			instructions = append(instructions, instruction)
		default:
			numberBuffer = append(numberBuffer, c)
		}
	}
	if len(numberBuffer) > 0 {
		flushBuffer(&instructions, &numberBuffer)
	}
	return instructions
}
