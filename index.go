package piscine

//Index is a good
func Index(s string, toFind string) int {
	if s == "" && toFind == "" {
		return 0
	}

	if getLen(s) < getLen(toFind) {
		return -1
	}

	for index := 0; index < getLen(s); index++ {
		if checkFrom(index, s, toFind) {
			return index
		}
	}

	return -1
}

func getLen(s string) int {
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

func checkFrom(start int, checkOn string, checkWith string) bool {

	if start+getLen(checkWith) > getLen(checkOn) {
		return false
	}

	runesCheckOn := []rune(checkOn)
	runesCheckWith := []rune(checkWith)
	for index := 0; index < getLen(checkWith); index++ {
		if runesCheckWith[index] != runesCheckOn[index+start] {
			return false
		}
	}

	return true
}
