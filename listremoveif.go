package piscine

//ListRemoveIf is a function
func ListRemoveIf(l *List, data_ref interface{}) {

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		return
	}

	current := l.Head.Next
	parent := l.Head
	for current != nil {
		if current.Data == data_ref {
			parent.Next = current.Next
			current = current.Next
			continue
		}
		parent = parent.Next
		current = current.Next
	}
}
