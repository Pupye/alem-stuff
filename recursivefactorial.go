package piscine

//RecursiveFactorial is a function that calculates factorial recursively
func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	if nb*RecursiveFactorial(nb-1) < 0 {
		return 0
	}
	return nb * RecursiveFactorial(nb-1)
}
