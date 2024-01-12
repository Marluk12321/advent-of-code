package movement

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

func moveLink(link *Link, cycleLength int) {
	prev := link.prev
	next := link.next
	connect(prev, next)

	fullCircleSteps := cycleLength - 1
	moveBy := link.value % fullCircleSteps
	var moveLeft, moveRight int
	if moveBy < 0 {
		moveLeft = moveBy
		moveRight = moveBy + fullCircleSteps
	} else {
		moveLeft = moveBy - fullCircleSteps
		moveRight = moveBy
	}

	if -moveLeft < moveRight {
		for ; moveLeft < 0; moveLeft++ {
			prev = prev.prev
		}
		next = prev.next
	} else {
		for ; moveRight > 0; moveRight-- {
			next = next.next
		}
		prev = next.prev
	}
	connect(prev, link)
	connect(link, next)
}

func toSlice(start *Link) []int {
	slice := []int{}
	link := start
	for {
		slice = append(slice, link.value)
		if link.next == start {
			break
		}
		link = link.next
	}
	return slice
}

func Move(numbers []int) []int {
	links := toLinks(numbers)
	connectIntoCycle(links)
	cycleLength := len(links)
	var zeroLink *Link
	for _, link := range links {
		moveLink(link, cycleLength)
		if link.value == 0 {
			zeroLink = link
		}
	}
	return toSlice(zeroLink)
}
