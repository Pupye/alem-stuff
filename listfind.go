package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head

	for current.Next != nil {
		if comp(ref, current.Data) {
			return &current.Data
		}
		current = current.Next
	}

	if comp(ref, current.Data) {
		return &current.Data
	}
	return nil
}
