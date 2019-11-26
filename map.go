package piscine

//Map is a good function
func Map(f func(int) bool, arr []int) []bool {
	length := getLenMap(arr)
	result := make([]bool, length)

	for index := range arr {
		result[index] = f(arr[index])
	}
	return result
}

func getLenMap(arr []int) int {
	length := 0
	isEmpty := true
	for x := range arr {
		length = x
		isEmpty = false
	}
	if !isEmpty {
		length++
	}
	return length
}
