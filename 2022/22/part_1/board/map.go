package board

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
	tiles  []Tile
}

func (row Row) IsWithinBounds(absoluteCol int) bool {
	rowMin := row.Offset
	rowMax := row.Offset + len(row.tiles) - 1
	return absoluteCol >= rowMin && absoluteCol <= rowMax
}

func (row Row) String() string {
	rowLen := len(row.tiles) + row.Offset
	representation := make([]rune, rowLen)
	for i := 0; i < row.Offset; i++ {
		representation[i] = ' '
	}
	for i, tile := range row.tiles {
		representation[i+row.Offset] = tile.Rune()
	}
	return string(representation)
}

type Map []Row

func BuildRow(text string) Row {
	row := Row{
		Offset: 0,
		tiles:  []Tile{},
	}
	for _, c := range text {
		switch c {
		case ' ':
			row.Offset++
		case '.':
			row.tiles = append(row.tiles, OPEN)
		case '#':
			row.tiles = append(row.tiles, WALL)
		default:
			panic(text)
		}
	}
	return row
}
