package terrain

type Position struct {
	Row int
	Col int
}

func absSub(x int, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func (position Position) DistanceTo(other Position) int {
	rowDist := absSub(position.Row, other.Row)
	colDist := absSub(position.Col, other.Col)
	return rowDist + colDist
}

func (position Position) Eq(other Position) bool {
	return position.Row == other.Row && position.Col == other.Col
}

type Terrain struct {
	HeightMap [][]rune
	Start     Position
	End       Position
}

func (terrain *Terrain) AddRow(rowDesc string) {
	row := []rune(rowDesc)
	for i, char := range row {
		switch char {
		case 'S':
			terrain.Start.Row = len(terrain.HeightMap)
			terrain.Start.Col = i
			row[i] = 'a'
		case 'E':
			terrain.End.Row = len(terrain.HeightMap)
			terrain.End.Col = i
			row[i] = 'z'
		}
	}
	terrain.HeightMap = append(terrain.HeightMap, row)
}

func (terrain *Terrain) GetHeight(position Position) rune {
	return terrain.HeightMap[position.Row][position.Col]
}

func (terrain *Terrain) GetRows() int {
	return len(terrain.HeightMap)
}

func (terrain *Terrain) GetCols() int {
	return len(terrain.HeightMap[0])
}
