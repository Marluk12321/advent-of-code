package system

type System struct {
	CPU   CPU
	CRT   CRT
	Cycle int
}

func MakeSystem() System {
	return System{
		CPU: MakeCPU(),
	}
}

func drawPixel(cpu *CPU, crt *CRT) {
	var pixel rune
	if cpu.x >= crt.currentCol-1 && cpu.x <= crt.currentCol+1 {
		pixel = '#'
	} else {
		pixel = '.'
	}
	crt.drawPixel(pixel)
}

func (system *System) StartCycle() {
	system.Cycle++
	drawPixel(&system.CPU, &system.CRT)
}

func (system *System) EndCycle() {
	system.CPU.endCycle()
	system.CRT.endCycle()
}
