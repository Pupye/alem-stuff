package piscine

func ListMerge(l1 *List, l2 *List) {
	if l2 == nil {
		return
	}

	if l2.Head == nil {
		return
	}

	if l1 == nil {
		return
	}

	if l1.Head == nil {
		l1.Head = l2.Head
		return
	}

	current := l1.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = l2.Head
	return
}
