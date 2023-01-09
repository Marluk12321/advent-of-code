package graph

import "2022/16/part_2/valve"

func makeFlowRates(valves []valve.Valve) map[string]int {
	flowRates := make(map[string]int, len(valves))
	for i := range valves {
		valve := &valves[i]
		flowRates[valve.Name] = valve.FlowRate
	}
	return flowRates
}

func makeDistances(valves []valve.Valve) map[[2]string]int {
	edgeCount := (len(valves) * (len(valves) - 1)) / 2
	distances := make(map[[2]string]int, edgeCount)
	return distances
}

func makeNeighborhood(valves []valve.Valve) map[string][]string {
	neighborhood := make(map[string][]string, len(valves))
	for i := range valves {
		valve := &valves[i]
		neighborhood[valve.Name] = valve.Neighbors
	}
	return neighborhood
}

func calcDistances(graph *Graph, neighborhood map[string][]string) {
	for startNode := range graph.FlowRates {
		front := []string{startNode}
		frontGeneration := 0
		seen := make(map[string]bool, len(graph.FlowRates))
		seen[startNode] = true
		for len(front) > 0 {
			var nextFront []string
			for _, node := range front {
				if frontGeneration > 0 {
					key := toDistanceKey(startNode, node)
					graph.distances[key] = frontGeneration
				}
				for _, neighbor := range neighborhood[node] {
					if !seen[neighbor] {
						nextFront = append(nextFront, neighbor)
						seen[neighbor] = true
					}
				}
			}
			front = nextFront
			frontGeneration++
		}
	}
}

func MakeGraph(valves []valve.Valve) Graph {
	graph := Graph{
		FlowRates: makeFlowRates(valves),
		distances: makeDistances(valves),
	}
	neighborhood := makeNeighborhood(valves)
	calcDistances(&graph, neighborhood)
	return graph
}
