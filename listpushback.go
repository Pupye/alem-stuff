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
	if l.Head == nil {
		l.Head = n
	} else {
		pointerToCurrent := l.Head
		for pointerToCurrent.Next != nil {
			pointerToCurrent = pointerToCurrent.Next
		}
		pointerToCurrent.Next = n
	}

}
