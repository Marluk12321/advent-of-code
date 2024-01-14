package tree

type Node struct {
	child1 *Node
	child2 *Node
	operation Operation
}

func (node Node) GetValue() int {
	return node.operation(node.child1, node.child2)
}

func toNode(desc NodeDesc) Node {
	var operation Operation
	switch desc.nodeType {
	case VALUE:
		operation = buildValueOperation(desc.value)
	case MATH:
		operation = buildMathOperation(desc.operator)
	default:
		panic(desc.nodeType)
	}
	return Node{operation: operation}
}

func toNodes(nodeDescs map[string]NodeDesc) map[string]*Node {
	nodes := make(map[string]*Node, len(nodeDescs))
	for _, desc := range nodeDescs {
		node := toNode(desc)
		nodes[desc.NodeName] = &node
	}
	return nodes
}

func connectIntoTree(
	nodes map[string]*Node,
	nodeDescs map[string]NodeDesc,
	rootName string,
) *Node {
	root := nodes[rootName]
	rootDesc := nodeDescs[rootName]
	if rootDesc.nodeType == MATH {
		root.child1 = connectIntoTree(nodes, nodeDescs, rootDesc.childName1)
		root.child2 = connectIntoTree(nodes, nodeDescs, rootDesc.childName2)
	}
	return root
}

func BuildTree(nodeDescs map[string]NodeDesc, rootName string) *Node {	
	nodes := toNodes(nodeDescs)
	return connectIntoTree(nodes, nodeDescs, rootName)
}
