package search

import (
	"2022/12/part_1/terrain"
	"container/heap"
	"fmt"
)

func getNeighbors(t *terrain.Terrain, position terrain.Position) []terrain.Position {
	neighbors := make([]terrain.Position, 0, 4)
	if position.Row > 0 {
		neighbors = append(neighbors, terrain.Position{
			Row: position.Row - 1,
			Col: position.Col,
		})
	}
	if position.Row < t.GetRows()-1 {
		neighbors = append(neighbors, terrain.Position{
			Row: position.Row + 1,
			Col: position.Col,
		})
	}
	if position.Col > 0 {
		neighbors = append(neighbors, terrain.Position{
			Row: position.Row,
			Col: position.Col - 1,
		})
	}
	if position.Col < t.GetCols()-1 {
		neighbors = append(neighbors, terrain.Position{
			Row: position.Row,
			Col: position.Col + 1,
		})
	}
	return neighbors
}

func getNext(t *terrain.Terrain, position terrain.Position) []terrain.Position {
	var next []terrain.Position
	currentHeight := t.GetHeight(position)
	for _, neighbor := range getNeighbors(t, position) {
		neighborHeight := t.GetHeight(neighbor)
		if neighborHeight <= currentHeight+1 {
			next = append(next, neighbor)
		}
	}
	return next
}

func AStar(t *terrain.Terrain) []terrain.Position {
	queue := makePathQueue(t)
	startPath := makePath(t.Start)
	heap.Push(&queue, startPath)
	bestSeen := map[terrain.Position]Path{}
	bestSeen[startPath.last] = startPath

	var positionsTested int
	for queue.Len() > 0 {
		obj := heap.Pop(&queue)
		path := obj.(Path)
		if path.last.Eq(t.End) {
			totalPositions := t.GetRows() * t.GetCols()
			fmt.Println("Tested:", positionsTested, "Total:", totalPositions, "Length:", path.len)
			return path.toForward()
		}
		positionsTested++

		for _, next := range getNext(t, path.last) {
			nextPath := path.join(next)
			previousBest, present := bestSeen[next]
			if !present || previousBest.len > nextPath.len {
				heap.Push(&queue, nextPath)
				bestSeen[next] = nextPath
			}
		}
	}

	panic("Path not found")
}

func BFS(t *terrain.Terrain) []terrain.Position {
	var queue []Path
	startPath := makePath(t.Start)
	queue = append(queue, startPath)
	seen := map[terrain.Position]bool{}
	seen[t.Start] = true

	var positionsTested int
	for len(queue) > 0 {
		path := &queue[0]
		queue = queue[1:]
		if path.last.Eq(t.End) {
			totalPositions := t.GetRows() * t.GetCols()
			fmt.Println("Tested:", positionsTested, "Total:", totalPositions)
			return path.toForward()
		}
		positionsTested++

		for _, next := range getNext(t, path.last) {
			if !seen[next] {
				nextPath := path.join(next)
				queue = append(queue, nextPath)
				seen[next] = true
			}
		}
	}

	panic("Path not found")
}
