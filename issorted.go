package piscine

//IsSorted is a function
func IsSorted(f func(a, b int) int, tab []int) bool {
	length := getLenIsSortedF(tab)

	for index := 0; index < length-1; index++ {
		if f(tab[index], tab[index+1]) > 0 {
			return false
		}
	}
	return true
}

func getLenIsSortedF(tab []int) int {
	isEmpty := true
	length := 0
	for index := range tab {
		length = index
		isEmpty = false
	}
	if !isEmpty {
		length++
	}
	return length
}
