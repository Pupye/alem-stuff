package piscine

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
