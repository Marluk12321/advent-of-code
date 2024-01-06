package objects

type Robot struct {
	Resource Resource
}

var (
	GEODE_CRACKER = Robot{Resource: OPEN_GEODE}
	OBSIDIAN_COLLECTOR = Robot{Resource: OBSIDIAN}
	CLAY_COLLECTOR = Robot{Resource: CLAY}
	ORE_COLLECTOR = Robot{Resource: ORE}
)

var robotForResource = map[Resource]Robot {
	OPEN_GEODE: GEODE_CRACKER,
	OBSIDIAN: OBSIDIAN_COLLECTOR,
	CLAY: CLAY_COLLECTOR,
	ORE: ORE_COLLECTOR,
}

var robotToName = map[Robot]string{
	GEODE_CRACKER: "geode-cracking robot",
	OBSIDIAN_COLLECTOR: "obsidian-collecting robot",
	CLAY_COLLECTOR: "clay-collecting robot",
	ORE_COLLECTOR: "ore-collecting robot",
}

func (robot Robot) String() string {
	name, found := robotToName[robot]
	if !found {
		return "unknown robot"
	}
	return name
}
