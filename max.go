package piscine

//Max is a function
func Max(arr []int) int {
	length := getLenOfArrMax(arr)
	if length == 0 {
		return 0 //here nay me error
	}

	if length == 1 {
		return arr[0]
	}

	max := arr[0]
	for index := range arr {
		if arr[index] > max {
			max = arr[index]
		}
	}
	return max
}

func getLenOfArrMax(arr []int) int {
	length := 0
	isEmpty := true

	for index := range arr {
		isEmpty = false
		length = index
	}
	if !isEmpty {
		length++
	}
	return length
}
