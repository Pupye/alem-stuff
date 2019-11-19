package piscine

//IsUpper is a function
func IsUpper(str string) bool {
	if str == "" {
		return false
	}

	mutableOne := []rune(str)
	for index := 0; index < getLenIsNLU(str); index++ {
		isUpper := mutableOne[index] > 64 && mutableOne[index] < 91
		if !(isUpper) {
			return false
		}
	}
	return true
}

func getLenIsNLU(s string) int {
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
