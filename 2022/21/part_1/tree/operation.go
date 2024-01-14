package tree

type Operator func(int, int) int

var (
	ADD Operator = func(x, y int) int {return x + y}
	SUB Operator = func(x, y int) int {return x - y}
	MUL Operator = func(x, y int) int {return x * y}
	DIV Operator = func(x, y int) int {return x / y}
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

type Operation func(*Node, *Node) int

func buildValueOperation(value int) Operation {
	return func(node1, node2 *Node) int {
		return value
	}
}

func buildMathOperation(operator Operator) Operation {
	return func(node1, node2 *Node) int {
		value1 := node1.GetValue()
		value2 := node2.GetValue()
		return operator(value1, value2)
	}
}
