package mixing

type Link struct {
	value int
	prev *Link
	next *Link
}

func toLinks(numbers []int) []*Link {
	links := make([]*Link, len(numbers))
	for i, number := range numbers {
		links[i] = &Link{value: number}
	}
	return links
}

func connect(prevLink, nextLink *Link) {
	prevLink.next = nextLink
	nextLink.prev = prevLink
}

func connectIntoCycle(links []*Link) {
	for i, link := range links[:len(links) - 1] {
		next := links[i + 1]
		connect(link, next)
	}
	first := links[0]
	last := links[len(links) - 1]
	connect(last, first)
}

func ToCycleLinks(numbers []int) []*Link {
	links := toLinks(numbers)
	connectIntoCycle(links)
	return links
}

func findZeroLink(links []*Link) *Link {
	for _, link := range links {
		if link.value == 0 {
			return link
		}
	}
	panic("no zero")
}

func ToZeroStartingSlice(links []*Link) []int {
	slice := []int{}
	zero := findZeroLink(links)
	link := zero
	for {
		slice = append(slice, link.value)
		if link.next == zero {
			break
		}
		link = link.next
	}
	return slice
}
