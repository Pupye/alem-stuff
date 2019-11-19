package piscine

//ToUpper is a function
func ToUpper(s string) string {
	mutableOne := []rune(s)

	for index := 0; index < len(mutableOne); index++ {
		if mutableOne[index] > 96 && mutableOne[index] < 123 {
			mutableOne[index] = mutableOne[index] - 32
		}
	}
	return string(mutableOne)
}
