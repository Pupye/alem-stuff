package piscine

//Unmatch is function
func Unmatch(arr []int) int {
	sortIntegerTableUNUNUN(arr)
	length := getLenUnmatch(arr)
	if length == 0 {
		return -1
	}
	if length == 1 {
		return arr[0]
	}
	for index := 0; index < length-1; index += 2 {
		if arr[index] != arr[index+1] {
			return arr[index]
		}
	}
	if length%2 == 1 {
		return arr[length-1]
	}
	return -1
}

func sortIntegerTableUNUNUN(table []int) { //validated
	key := 0
	j := 0
	length := getLenUnmatch(table)
	for index := 1; index < length; index++ {
		key = table[index]
		j = index - 1

		for j >= 0 && table[j] > key {
			table[j+1] = table[j]
			j = j - 1
		}
		table[j+1] = key
	}
}

func getLenUnmatch(s []int) int { //validated

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
