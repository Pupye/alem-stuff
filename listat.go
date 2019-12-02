package piscine

//NodeL is a struct
type NodeL struct {
	Data interface{}
	Next *NodeL
}

//ListAt is function
func ListAt(l *NodeL, pos int) *NodeL {
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
