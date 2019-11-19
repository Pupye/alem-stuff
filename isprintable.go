package piscine

//IsPrintable is a function
func IsPrintable(str string) bool {
	if str == "" {
		return false
	}

	mutableOne := []rune(str)
	for index := 0; index < getLenIsNPQ(str); index++ {
		isPrintable := mutableOne[index] > 31 && mutableOne[index] < 128
		if !(isPrintable) {
			return false
		}
	}
	return true
}

func getLenIsNPQ(s string) int {
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
