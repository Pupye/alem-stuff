package piscine

func IsPositive_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsNegative_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsNotNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	case string, rune:
		return true
	}
	return true
}

//ListForEachIf is function
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l.Head == nil {
		return
	}

	current := l.Head
	for current.Next != nil {
		if cond(current) {
			f(current)
		}

		current = current.Next
	}
	if cond(current) {
		f(current)
	}
}
