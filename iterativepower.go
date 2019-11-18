package piscine

//IterativePower is a function that calculates power
func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	result := 1
	for index := 0; index < power; index++ {
		result *= nb
	}
	return result
}
