package piscine

//AtoiBase is a function
func AtoiBase(s string, base string) int {
	if !isValidBaseForAtoiBase(base) {
		return 0
	}

	if !isStringInBase(s, base) {
		return 0
	}

	length := getLenOfAtoiBase(s)
	radix := getLenOfAtoiBase(base)
	baseRad := 1
	result := 0
	mutableOne := []rune(s)
	for index := length - 1; index > -1; index-- {
		digit := getIndex(base, mutableOne[index])
		result += digit * baseRad
		baseRad *= radix
	}
	return result
}

func getLenOfAtoiBase(str string) int { //validated
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

func getIndex(base string, char rune) int {
	count := 0
	for _, v := range base {
		if char == v {
			return count
		}
		count++
	}
	return count
}

func isValidBaseForAtoiBase(base string) bool {
	length := getLenOfStrBase(base)

	if length < 2 {
		return false
	}
	mutableOne := []rune(base)

	for x := range mutableOne {
		if mutableOne[x] == '+' || mutableOne[x] == '-' {
			return false
		}
		for index := x + 1; index < length; index++ {

			if mutableOne[x] == mutableOne[index] {

				return false
			}
		}
	}

	return true
}

func isStringInBase(s, base string) bool {
	for _, v := range s {
		if !inBase(v, base) {
			return false
		}
	}
	return true
}

func inBase(v rune, base string) bool {
	for _, char := range base {
		if char == v {
			return true
		}
	}
	return false
}
