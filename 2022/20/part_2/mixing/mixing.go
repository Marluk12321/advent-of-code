package mixing

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

func Mix(links []*Link) {
	cycleLength := len(links)
	for _, link := range links {
		moveLink(link, cycleLength)
	}
}
