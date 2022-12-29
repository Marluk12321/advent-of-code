package monkeys

type Monkeys struct {
	List         []Monkey
	Current      int
	PlayedRounds int
}

func (monkeys *Monkeys) Add(monkey *Monkey) {
	monkeys.List = append(monkeys.List, *monkey)
}

func (monkeys *Monkeys) TakeTurn() {
	monkey := &monkeys.List[monkeys.Current]
	for monkey.Items.anyLeft() {
		monkey.inspect()
		monkey.loseInterest()
		targetIndex := monkey.pointNext()
		targetMonkey := &monkeys.List[targetIndex]
		monkey.passItem(targetMonkey)
	}
	monkeys.Current++
	if monkeys.Current == len(monkeys.List) {
		monkeys.Current = 0
		monkeys.PlayedRounds++
	}
}
