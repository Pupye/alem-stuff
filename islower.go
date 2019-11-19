package piscine

//IsLower is a function
func IsLower(str string) bool {
	if str == "" {
		return false
	}

	mutableOne := []rune(str)
	for index := 0; index < getLenIsNL(str); index++ {
		isLower := mutableOne[index] > 96 && mutableOne[index] < 123
		if !(isLower) {
			return false
		}
	}
	return true
}

func getLenIsNL(s string) int {
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
