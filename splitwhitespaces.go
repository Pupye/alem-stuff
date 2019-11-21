package piscine

//SplitWhiteSpaces is a function
func SplitWhiteSpaces(str string) []string {
	numOfWords := 0
	for x := range str {
		if str[x] == ' ' {
			numOfWords++
		}
	}
	result := make([]string, numOfWords)

	tempStr := ""
	index := 0
	for x := range str {
		tempStr += string(str[x])
		if str[x] == ' ' {
			result[index] = tempStr
			index++
			tempStr = ""
		}
	}
	return result
}
