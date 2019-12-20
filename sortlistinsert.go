package piscine

//NodeI is a struct for this execise
type NodeI struct {
	Data int
	Next *NodeI
}

//SortListInsert is a function that inserts data keeping the order
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	if l == nil {
		return n
	}

	if n.Data < l.Data {
		n.Next = l
		return n
	}

	parent := l
	child := l.Next
	for child != nil && n.Data > child.Data {
		parent = parent.Next
		child = child.Next
	}

	if child == nil {
		parent.Next = n
	} else {
		n.Next = child
		parent.Next = n
	}
	return l
}
