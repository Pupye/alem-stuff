package piscine

//CountIf is a function
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for index := range tab {
		if f(tab[index]) {
			count++
		}
	}
	return count
}
