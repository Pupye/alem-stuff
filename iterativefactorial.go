package piscine

//IterativeFactorial is a function that calculates factorial itratively
func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for index := 1; index <= nb; index++ {
		temp := result
		result = temp * index
		if result/temp != index {
			return 0
		}
	}

	return result
}
