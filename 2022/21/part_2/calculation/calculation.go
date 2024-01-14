package calculation

import "2022/21/part_2/tree"

func calcNextTargetValue(
	operator tree.Operator,
	varSide tree.Direction,
	resultValue,
	otherValue int,
) int {
	switch operator {
	case tree.ADD:
		return resultValue - otherValue
	case tree.SUB:
		switch varSide {
		case tree.LEFT:
			return resultValue + otherValue
		case tree.RIGHT:
			return otherValue - resultValue
		default:
			panic(varSide)
		}
	case tree.MUL:
		return resultValue / otherValue
	case tree.DIV:
		switch varSide {
		case tree.LEFT:
			return resultValue * otherValue
		case tree.RIGHT:
			return otherValue / resultValue
		default:
			panic(varSide)
		}
	default:
		panic(operator)
	}
}

func matchTargetValue(
	nodeDescs map[string]tree.NodeDesc,
	rootName string,
	varPath tree.Path,
	targetValue int,
) int {
	if len(varPath) == 0 {
		return targetValue
	}
	rootDesc := nodeDescs[rootName]
	varSide := varPath.Head()
	varSideName := varSide.Get(rootDesc)
	oppositeSideName := varSide.Opposite().Get(rootDesc)
	oppositeValue := tree.BuildTree(nodeDescs, oppositeSideName).GetValue()
	nextTargetValue := calcNextTargetValue(rootDesc.Operator, varSide, targetValue, oppositeValue)
	return matchTargetValue(nodeDescs, varSideName, varPath.Tail(), nextTargetValue)
}

func CalculateVariableValue(
	nodeDescs map[string]tree.NodeDesc,
	rootName, varName string,
) int {
	varPath, found := tree.FindPath(nodeDescs, rootName, varName)
	if !found {
		panic("variable not found")
	}
	rootDesc := nodeDescs[rootName]
	varSide := varPath.Head()
	varSideName := varSide.Get(rootDesc)
	oppositeSideName := varSide.Opposite().Get(rootDesc)
	oppositeValue := tree.BuildTree(nodeDescs, oppositeSideName).GetValue()
	return matchTargetValue(nodeDescs, varSideName, varPath.Tail(), oppositeValue)
}
