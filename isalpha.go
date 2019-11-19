package piscine

//IsAlpha is a function
func IsAlpha(str string) bool {
	if str == "" {
		return false
	}

	mutableOne := []rune(str)
	for index := 0; index < getLenIsA(str); index++ {
		isNumeric := mutableOne[index] > 47 && mutableOne[index] < 58
		isAlpha := (mutableOne[index] > 64 && mutableOne[index] < 91) || (mutableOne[index] > 96 && mutableOne[index] < 123)
		if !(isAlpha || isNumeric) {
			return false
		}
	}
	return true
}

func getLenIsA(s string) int {
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
