package piscine

//SortedListMerge is a function to merge two different lists
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	res := &NodeI{}
	current := res
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			current.Next = &NodeI{Data: n1.Data}
			current = current.Next
			n1 = n1.Next
		} else {
			current.Next = &NodeI{Data: n2.Data}
			current = current.Next
			n2 = n2.Next
		}
	}
	if n1 == nil {
		for n2 != nil {
			current.Next = &NodeI{Data: n2.Data}
			current = current.Next
			n2 = n2.Next
		}
	} else {
		for n1 != nil {
			current.Next = &NodeI{Data: n1.Data}
			current = current.Next
			n1 = n1.Next
		}
	}
	return res.Next
}
