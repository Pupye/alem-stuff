package piscine

//ListRemoveIf is a function
func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	current := l.Head

	if current == nil {
		return
	}

	for current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		}
		current = current.Next
	}

	if current.Data == data_ref { //takes care of last node
		current = nil
	}

}
