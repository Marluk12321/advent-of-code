package search

import "2022/12/part_2/terrain"

type Path struct {
	prev *Path
	last terrain.Position
	len  int
}

func makePath(position terrain.Position) Path {
	return Path{
		last: position,
		len:  1,
	}
}

func (path *Path) join(position terrain.Position) Path {
	return Path{
		prev: path,
		last: position,
		len:  path.len + 1,
	}
}

func (path *Path) toForward() []terrain.Position {
	result := make([]terrain.Position, path.len)
	i := len(result) - 1
	for path != nil {
		result[i] = path.last
		path = path.prev
		i--
	}
	return result
}
