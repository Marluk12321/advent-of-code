package mapping

type Tile int

const (
	OPEN Tile = iota
	WALL
)

func (tile Tile) Rune() rune {
	switch tile {
	case OPEN:
		return '.'
	case WALL:
		return '#'
	default:
		panic(tile)
	}
}

type Row struct {
	Offset int
	Tiles  []Tile
}

func (row Row) IsWithinBounds(absoluteCol int) bool {
	rowMin := row.Offset
	rowMax := row.Offset + len(row.Tiles) - 1
	return absoluteCol >= rowMin && absoluteCol <= rowMax
}

func (row Row) String() string {
	rowLen := len(row.Tiles) + row.Offset
	representation := make([]rune, rowLen)
	for i := 0; i < row.Offset; i++ {
		representation[i] = ' '
	}
	for i, tile := range row.Tiles {
		representation[i+row.Offset] = tile.Rune()
	}
	return string(representation)
}

type Board []Row

func BuildRow(text string) Row {
	row := Row{
		Offset: 0,
		Tiles:  []Tile{},
	}
	for _, c := range text {
		switch c {
		case ' ':
			row.Offset++
		case '.':
			row.Tiles = append(row.Tiles, OPEN)
		case '#':
			row.Tiles = append(row.Tiles, WALL)
		default:
			panic(text)
		}
	}
	return row
}
