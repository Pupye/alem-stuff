package piscine

//List is a struct
type List struct {
	Head *NodeL
	Tail *NodeL
}

//ListLast is a function that returns last element of the list
func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	pointerToCurrnet := l.Head

	for pointerToCurrnet.Next != nil {
		pointerToCurrnet = pointerToCurrnet.Next
	}

	return pointerToCurrnet.Data
}
