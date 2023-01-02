package search

import "2022/12/part_1/terrain"

type PathPriority struct {
	path             Path
	distanceToTarget int
}

type PathQueue struct {
	storage []PathPriority
	target  terrain.Position
}

func makePathQueue(t *terrain.Terrain) PathQueue {
	return PathQueue{
		target: t.End,
	}
}

func (queue *PathQueue) Len() int {
	return len(queue.storage)
}

func (queue *PathQueue) Less(i int, j int) bool {
	entryI := &queue.storage[i]
	entryJ := &queue.storage[j]
	distanceI := entryI.path.len + entryI.distanceToTarget
	distanceJ := entryJ.path.len + entryJ.distanceToTarget
	return distanceI < distanceJ
}

func (queue *PathQueue) Swap(i int, j int) {
	storage := &queue.storage
	(*storage)[i], (*storage)[j] = (*storage)[j], (*storage)[i]
}

func (queue *PathQueue) Push(obj any) {
	path := obj.(Path)
	entry := PathPriority{
		path:             path,
		distanceToTarget: path.last.DistanceTo(queue.target),
	}
	queue.storage = append(queue.storage, entry)
}

func (queue *PathQueue) Pop() any {
	index := queue.Len() - 1
	entry := queue.storage[index]
	queue.storage = queue.storage[:index]
	return entry.path
}
