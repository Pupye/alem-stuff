package piscine

//SortIntegerTable is a function that sorts in arr
func SortIntegerTable(table []int) {
	key := 0
	j := 0
	for index := 1; index < len(table); index++ {
		key = table[index]
		j = index - 1

		for j >= 0 && table[j] > key {
			table[j+1] = table[j]
			j = j - 1
		}
		table[j+1] = key
	}
}
