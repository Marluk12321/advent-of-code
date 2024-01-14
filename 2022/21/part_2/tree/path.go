package tree

type Direction int

const (
	LEFT Direction = iota
	RIGHT
)

func (direction Direction) Opposite() Direction {
	switch direction {
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	default:
		panic(direction)
	}
}

func (direction Direction) Get(desc NodeDesc) string {
	switch direction {
	case LEFT:
		return desc.childName1
	case RIGHT:
		return desc.childName2
	default:
		panic(direction)
	}
}

type Path []Direction

func (path Path) Head() Direction {
	return path[0]
}

func (path Path) Tail() Path {
	return path[1:]
}

func FindPath(nodeDescs map[string]NodeDesc, rootName, targetName string) (Path, bool) {
	path := []Direction{}
	desc := nodeDescs[rootName]
	if desc.NodeName == targetName {
		return path, true
	}
	if desc.nodeType == VALUE {
		return path, false
	}
	if subPath, found := FindPath(nodeDescs, desc.childName1, targetName); found {
		path = append(path, LEFT)
		path = append(path, subPath...)
		return path, found
	}
	if subPath, found := FindPath(nodeDescs, desc.childName2, targetName); found {
		path = append(path, RIGHT)
		path = append(path, subPath...)
		return path, found
	}
	return path, false
}
