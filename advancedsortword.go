package piscine

//AdvancedSortWordArr is function that sorts
func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	key := ""
	j := 0
	length := getLenAdvancedSort(array)
	for index := 1; index < length; index++ {
		key = array[index]
		j = index - 1
		for j >= 0 && f(array[j], key) > 0 {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
}

func getLenAdvancedSort(s []string) int { //validated

	length := 0
	isEmpty := true
	for x := range s {
		length = x
		isEmpty = false
	}
	if !isEmpty {
		length++
	}

	return length
}
