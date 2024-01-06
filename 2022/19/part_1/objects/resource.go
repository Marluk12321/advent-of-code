package objects

type Resource int

const (
	OPEN_GEODE Resource = iota
	OBSIDIAN
	CLAY
	ORE
)

var resourceFromName = map[string]Resource {
	"geode": OPEN_GEODE,
	"obsidian": OBSIDIAN,
	"clay": CLAY,
	"ore": ORE,
}

var resourceToName = map[Resource]string {
	OPEN_GEODE: "geode",
	OBSIDIAN: "obsidian",
	CLAY: "clay",
	ORE: "ore",
}

func (resource Resource) String() string {
	name, found := resourceToName[resource]
	if !found {
		return "unknown resource"
	}
	return name
}

type Resources map[Resource]int
