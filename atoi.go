package piscine

//Atoi is function that converts string to integer
func Atoi(s string) int {
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
		if mutableOne[index] == 45 && index == 0 {
			return res * (-1)
		} else if mutableOne[index] == 43 && index == 0 {
			return res
		}
		if mutableOne[index] < 48 || mutableOne[index] > 57 {
			return 0
		}
		slicedInt := mutableOne[index] - 48
		res += int(slicedInt) * tens
		tens *= 10
	}

	return res
}
