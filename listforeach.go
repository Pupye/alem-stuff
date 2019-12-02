package piscine

//ListForEach is a function
func ListForEach(l *List, f func(*NodeL)) {
	if l.Head == nil {
		//do nothing
		return
	}

	current := l.Head
	for current.Next != nil {
		f(current)
		current = current.Next
	}
	f(current)
}

//Add2_node is also function
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

//Subtract3_node is also function
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
