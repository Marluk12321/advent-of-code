package valve

import (
	"strconv"
	"strings"
)

type Valve struct {
	Name      string
	FlowRate  int
	Neighbors []string
}

func toInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return value
}

func toValve(valveDefinition string) Valve {
	fields := strings.Fields(valveDefinition)
	name := fields[1]
	flowRateText := fields[4][len("text="):]
	flowRate := toInt(flowRateText)
	return Valve{
		Name:     name,
		FlowRate: flowRate,
	}
}

func fillNeightbors(valve *Valve, neighborsDefinition string) {
	prefixLen := len("tunnels lead to valve")
	switch neighborsDefinition[prefixLen] {
	case ' ':
		neighborsDefinition = neighborsDefinition[prefixLen+1:]
	case 's':
		neighborsDefinition = neighborsDefinition[prefixLen+2:]
	}
	neighbors := strings.Split(neighborsDefinition, ", ")
	valve.Neighbors = append(valve.Neighbors, neighbors...)
}

func MakeValve(text string) Valve {
	parts := strings.Split(text, "; ")
	valve := toValve(parts[0])
	fillNeightbors(&valve, parts[1])
	return valve
}
