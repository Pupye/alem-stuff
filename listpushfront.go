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

//ListPushFront pushes to the front
func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}
