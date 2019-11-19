package piscine

//Compare is a function
func Compare(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)

	checkUpTo := 0
	if getLenC(a) < getLenC(b) {
		checkUpTo = getLenC(a)
	} else {
		checkUpTo = getLenC(b)
	}
	for index := 0; index < checkUpTo; index++ {
		if runesA[index] < runesB[index] {
			return -1
		} else if runesA[index] > runesB[index] {
			return 1
		}
	}
	if getLenC(a) < getLenC(b) {
		return -1
	} else if getLenC(a) > getLenC(b) {
		return 1
	}

	return 0
}

func getLenC(s string) int {
	if s == "" {
		return 0
	}

	length := 0
	runes := []rune(s)
	for x := range runes {
		length = x
	}
	length++

	return length
}
