package checker

func compareInts(left int, right int) int {
	if left < right {
		return -1
	}
	if left == right {
		return 0
	}
	return 1
}

func toList(x int) []interface{} {
	return []interface{}{x}
}

func compareLists(left []interface{}, right []interface{}) int {
	lLen := len(left)
	rLen := len(right)
	var i int
	for {
		lEnd := i == lLen
		rEnd := i == rLen
		if lEnd {
			if rEnd {
				return 0
			} else {
				return -1
			}
		} else if rEnd {
			return 1
		}

		elemCmpResult := compare(&left[i], &right[i])
		if elemCmpResult != 0 {
			return elemCmpResult
		}
		i++
	}
}

func compare(left *interface{}, right *interface{}) int {
	switch (*left).(type) {
	case int:
		lValue := (*left).(int)
		switch (*right).(type) {
		case int:
			rValue := (*right).(int)
			return compareInts(lValue, rValue)
		case []interface{}:
			rValue := (*right).([]interface{})
			return compareLists(toList(lValue), rValue)
		default:
			panic(*right)
		}
	case []interface{}:
		lValue := (*left).([]interface{})
		switch (*right).(type) {
		case int:
			rValue := (*right).(int)
			return compareLists(lValue, toList(rValue))
		case []interface{}:
			rValue := (*right).([]interface{})
			return compareLists(lValue, rValue)
		default:
			panic(*right)
		}
	default:
		panic(*left)
	}
}

func InRightOrder(left *interface{}, right *interface{}) bool {
	return compare(left, right) < 0
}
