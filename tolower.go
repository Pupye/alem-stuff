package piscine

//ToLower is a function
func ToLower(s string) string {
	mutableOne := []rune(s)

	for index := 0; index < len(mutableOne); index++ {
		if mutableOne[index] > 64 && mutableOne[index] < 91 {
			mutableOne[index] = mutableOne[index] + 32
		}
	}
	return string(mutableOne)
}
