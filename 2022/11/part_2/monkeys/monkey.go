package monkeys

type Monkey struct {
	Items          ItemQueue
	updateWorry    func(int) int
	getTarget      func(int) int
	worryReduction int
}

func (monkey *Monkey) inspect() {
	itemWorry := monkey.Items.take()
	itemWorry = monkey.updateWorry(itemWorry)
	if itemWorry < 0 {
		panic(itemWorry)
	}
	monkey.Items.putBack(itemWorry)
}

func (monkey *Monkey) loseInterest() {
	itemWorry := monkey.Items.take()
	monkey.Items.putBack(itemWorry % monkey.worryReduction)
}

func (monkey *Monkey) pointNext() int {
	itemWorry := monkey.Items.first()
	return monkey.getTarget(itemWorry)
}

func (monkey *Monkey) passItem(other *Monkey) {
	item := monkey.Items.take()
	other.Items.add(item)
}
