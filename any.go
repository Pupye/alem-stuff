package piscine

//Any is a function
func Any(f func(string) bool, arr []string) bool {
	for index := range arr {
		if f(arr[index]) {
			return true
		}
	}
	return false
}
