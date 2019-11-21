package piscine

//SplitWhiteSpaces is a function
func SplitWhiteSpaces(str string) []string {
	numOfWords := 0
	encountered := false
	for x := range str {
		isDelimeter := (str[x] == ' ') || (str[x] == '\t') || (str[x] == '\n')
		if isDelimeter {
			encountered = false
			continue
		} else {
			if !encountered {
				numOfWords++
				encountered = true
			}

		}
	}

	result := make([]string, numOfWords)
	tempStr := ""
	index := 0
	strLen := getLenSplitSpace(str)

	for x := 0; x < strLen; x++ {
		isDelimeter := (str[x] == ' ') || (str[x] == '\t') || (str[x] == '\n')
		if isDelimeter {
			continue
		} else {

			for !isDelimeter {
				tempStr += string(str[x])

				x++
				if x == strLen {
					break
				}
				isDelimeter = (str[x] == ' ') || (str[x] == '\t') || (str[x] == '\n')
			}
			result[index] = tempStr
			tempStr = ""
			index++
		}
	}
	return result
}

func getLenSplitSpace(str string) int {
	length := 0
	isEmpty := true
	for index := range str {
		isEmpty = false
		length = index
	}
	if !isEmpty {
		length++
	}

	return length
}
