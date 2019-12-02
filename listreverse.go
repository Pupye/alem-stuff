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

//ListReverse is a function
func ListReverse(l *List) {
	size := listSizeForReverse(l)

	temp := &List{}

	for index := size - 1; index > -1; index-- {
		tempNode := listAtForReverse(l.Head, index)
		listPushBackForReverse(temp, tempNode.Data)
	}
	l.Head = nil
	l.Tail = nil

	for index := 0; index < size; index++ {
		tempNode := listAtForReverse(temp.Head, index)
		listPushBackForReverse(l, tempNode.Data)
	}
}

func listSizeForReverse(l *List) int {
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

func listAtForReverse(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	if l == nil {
		return nil
	}
	current := l
	index := 0
	for index < pos {
		if current.Next == nil {
			return nil
		}
		current = current.Next
		index++
	}
	return current
}

func listPushBackForReverse(l *List, data interface{}) {
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
