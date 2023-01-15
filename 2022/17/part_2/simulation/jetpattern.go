package simulation

type JetPattern struct {
	xOffests []int
	index    int
}

func (pattern *JetPattern) NextXOffset() int {
	offset := pattern.xOffests[pattern.index]
	pattern.index = (pattern.index + 1) % len(pattern.xOffests)
	return offset
}

func MakeJetPattern(patternDesc string) JetPattern {
	offsets := make([]int, len(patternDesc))
	for i, direction := range patternDesc {
		switch direction {
		case '<':
			offsets[i] = -1
		case '>':
			offsets[i] = 1
		default:
			panic(direction)
		}
	}
	return JetPattern{xOffests: offsets}
}
