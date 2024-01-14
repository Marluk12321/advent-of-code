package tree

import (
	"strconv"
	"strings"
)

type NodeType int

const (
	VALUE NodeType = iota
	MATH
)

type NodeDesc struct {
	NodeName   string
	nodeType   NodeType
	Operator   Operator
	childName1 string
	childName2 string
	value      int
}

func fillOperationDetails(desc *NodeDesc, text string) {
	text = strings.TrimSpace(text)
	parts := strings.Split(text, " ")
	if len(parts) == 1 {
		value, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		desc.nodeType = VALUE
		desc.value = value
	} else if len(parts) == 3 {
		desc.nodeType = MATH
		desc.childName1 = parts[0]
		desc.Operator = toOperator(parts[1])
		desc.childName2 = parts[2]
	} else {
		panic(parts)
	}
}

func BuildNodeDesc(text string) NodeDesc {
	parts := strings.Split(text, ": ")
	desc := NodeDesc{
		NodeName: parts[0],
	}
	fillOperationDetails(&desc, parts[1])
	return desc
}
