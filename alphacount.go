package piscine

//AlphaCount is a function that counts number of chars
func AlphaCount(str string) int {
	mutableOne := []byte(str)
	count := 0
	for x := range mutableOne {
		if (mutableOne[x] > 64 && mutableOne[x] < 91) || (mutableOne[x] > 96 && mutableOne[x] < 123) {
			count++
		}
	}
	return count
}
