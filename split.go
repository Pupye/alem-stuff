package piscine

func getLenOfSplit(str string) int { //validated
	isEmpty := true
	length := 0
	for index := range str {
		isEmpty = false
		length = index
	}
	if !isEmpty {
		length++
	}
	return length
}

//Split is a function
func Split(str, charset string) []string {
	numOfWord := getNumOfWordsSplit(str, charset)
	result := make([]string, numOfWord)
	wordBuilder := ""
	wordOrder := 0
	mtableStr := []rune(str)
	length := getLenOfSplit(str)

	for index := 0; index < length; index++ {
		if isCharsetPresent(mtableStr, index, []rune(charset)) {
			index += (getLenOfSplit(charset) - 1)
		} else {
			for !isCharsetPresent(mtableStr, index, []rune(charset)) {
				wordBuilder += string(mtableStr[index])
				index++
				if index == length {
					break
				}
			}
			index += (getLenOfSplit(charset) - 1)
			result[wordOrder] = wordBuilder
			wordBuilder = ""
			wordOrder++
		}
	}
	return result
}

func getNumOfWordsSplit(str, charset string) int { //validated
	encountered := false
	count := 0
	lengthStr := getLenOfSplit(str)
	for index := 0; index < lengthStr; index++ {
		if isCharsetPresent([]rune(str), index, []rune(charset)) {
			encountered = false
			index += (getLenOfSplit(charset) - 1)
			continue
		} else {
			if !encountered {
				count++
				encountered = true
			}
		}
	}
	return count
}

func isCharsetPresent(mtableStr []rune, start int, charset []rune) bool { //validated

	lengthCharSet := getLenOfSplit(string(charset))
	lengthStr := getLenOfSplit(string(mtableStr))

	if start+lengthCharSet > lengthStr {
		return false
	}

	for index := 0; index < lengthCharSet; index++ {
		if mtableStr[index+start] != charset[index] {
			return false
		}
	}
	return true
}
