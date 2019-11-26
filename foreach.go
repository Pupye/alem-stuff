package piscine

//ForEach is a function
func ForEach(f func(int), arr []int) {
	for index := range arr {
		f(arr[index])
	}
}
