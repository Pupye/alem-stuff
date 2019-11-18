package piscine

//Sqrt is a function that calculates sqrt of a number
func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb < 2 {
		return nb
	}
	for index := 2; index < nb; index++ {
		sqr := index * index
		if sqr == nb {
			return index
		}

		if sqr > nb {
			return 0
		}
	}
	return 0
}
