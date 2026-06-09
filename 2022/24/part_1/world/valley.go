package world

type Valley struct {
	Rows, Cols, EntryCol, ExitCol int
	Blizzards  []Blizzard
}

func MakeValley() Valley {
	return Valley{Rows: 0, Cols: 0, Blizzards: make([]Blizzard, 0)}
}

func (valley Valley) String() string {
	
}
