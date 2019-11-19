package piscine

//IsNumeric is a function
func IsNumeric(str string) bool {
	if str == "" {
		return false
	}

	mutableOne := []rune(str)
	for index := 0; index < getLenIsN(str); index++ {
		isNumeric := mutableOne[index] > 47 && mutableOne[index] < 58
		if !(isNumeric) {
			return false
		}
	}
	return true
}

func getLenIsN(s string) int {
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
