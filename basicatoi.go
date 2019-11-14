package piscine

//BasicAtoi is function that converts string to integer
func BasicAtoi(s string) int {
	if s == "" {
		return 0
	}
	mutableOne := []rune(s)
	strLen := 0
	for s := range mutableOne {
		strLen = s
	}
	strLen++ // now we got len
	tens := 1
	res := 0
	for index := strLen - 1; index >= 0; index-- {
		slicedInt := mutableOne[index] - 48
		res += int(slicedInt) * tens
		tens *= 10
	}

	return res
}
