package graph

type Graph struct {
	FlowRates map[string]int
	distances map[[2]string]int
}

func toDistanceKey(name1 string, name2 string) [2]string {
	if name1 < name2 {
		return [2]string{name1, name2}
	} else {
		return [2]string{name2, name1}
	}
}

func (graph *Graph) GetDistance(name1 string, name2 string) int {
	key := toDistanceKey(name1, name2)
	value, exists := graph.distances[key]
	if !exists {
		panic(key)
	}
	return value
}
