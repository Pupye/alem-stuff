package piscine

//RecursiveFactorial is a function that calculates factorial recursively
func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}

	if nb < 0 {
		return 0
	}
	temp := RecursiveFactorial(nb - 1)
	result := nb * temp
	if result/nb != temp {
		return 0
	}
	return result
}
