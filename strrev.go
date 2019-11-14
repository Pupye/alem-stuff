package piscine

//StrRev is function that reverses string
func StrRev(s string) string {
	strLen := 0
	mutableOne := []rune(s)
	for s := range mutableOne {
		strLen = s
	}
	strLen++ // now we got len
	startIndex := 0
	endIndex := strLen - 1
	for startIndex < endIndex {
		temp := mutableOne[startIndex]
		mutableOne[startIndex] = mutableOne[endIndex]
		mutableOne[endIndex] = temp
		startIndex++
		endIndex--
	}
	return string(mutableOne)
}
