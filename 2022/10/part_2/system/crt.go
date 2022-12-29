package system

import "fmt"

type CRT struct {
	buffer     [6][40]rune
	currentRow int
	currentCol int
}

func (crt *CRT) endCycle() {
	crt.currentCol++
	if crt.currentCol >= len(crt.buffer[crt.currentRow]) {
		crt.currentCol = 0
		crt.currentRow++
		if crt.currentRow >= len(crt.buffer) {
			crt.currentRow = 0
		}
	}
}

func (crt *CRT) drawPixel(value rune) {
	crt.buffer[crt.currentRow][crt.currentCol] = value
}

func (crt *CRT) Print() {
	for _, row := range crt.buffer {
		fmt.Println(string(row[:]))
	}
}
