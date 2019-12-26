package piscine

//SortWordArr is a function
func SortWordArr(array []string) {
	key := ""
	j := 0
	length := getLenAdvancedSort(array)
	for index := 1; index < length; index++ {
		key = array[index]
		j = index - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
}
