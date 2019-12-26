package piscine

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	res := &NodeI{Data: -9223372036854775808}
	for l != nil {
		res = SortListInsert(res, l.Data)
		l = l.Next
	}
	return res.Next
}
