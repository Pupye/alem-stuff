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

//ListPushBack is a function to push back the data
func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Tail == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = l.Tail.Next
	}
}
