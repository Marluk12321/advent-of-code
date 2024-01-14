package tree

type Operator rune

const (
	ADD Operator = '+'
	SUB Operator = '-'
	MUL Operator = '*'
	DIV Operator = '/'
)

var nameToOperator = map[string]Operator{
	"+": ADD,
	"-": SUB,
	"*": MUL,
	"/": DIV,
}

func toOperator(text string) Operator {
	operator, found := nameToOperator[text]
	if !found {
		panic(text)
	}
	return operator
}

type Operation func(*Node) int

func buildValueOperation(value int) Operation {
	return func(node *Node) int {
		return value
	}
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

var operatorToFunc = map[Operator]func(int, int) int {
	ADD: add,
	SUB: sub,
	MUL: mul,
	DIV: div,
}

func buildMathOperation(operator Operator) Operation {
	operatorFunc, found := operatorToFunc[operator]
	if !found {
		panic(operator)
	}
	return func(node *Node) int {
		value1 := node.child1.GetValue()
		value2 := node.child2.GetValue()
		return operatorFunc(value1, value2)
	}
}
