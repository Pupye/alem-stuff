package piscine

//NodeL is a struct
type NodeL struct {
	Data interface{}
	Next *NodeL
}

//List is a struct
type List struct {
	Head *NodeL
	Tail *NodeL
}

//ListSize function that return list size
func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	size := 1

	temp := l.Head

	for temp.Next != nil {
		size++
		temp = temp.Next
	}
	return size
}
