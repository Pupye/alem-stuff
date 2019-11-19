package piscine

//Capitalize is a function
func Capitalize(s string) string {
	if s == "" {
		return s
	}

	mutableOne := []rune(s)
	length := getLenCap(s)

	encountered := false
	for index := 0; index < length; index++ {

		if !encountered && (mutableOne[index] > 64 && mutableOne[index] < 91) {
			encountered = true
		} else if !encountered && (mutableOne[index] > 96 && mutableOne[index] < 123) {
			encountered = true
			mutableOne[index] -= 32
		} else if !encountered && (mutableOne[index] > 47 && mutableOne[index] < 58) {
			encountered = true
		} else if encountered == true {
			if mutableOne[index] > 64 && mutableOne[index] < 91 {
				mutableOne[index] += 32
			} else if !((mutableOne[index] > 96 && mutableOne[index] < 123) || (mutableOne[index] > 47 && mutableOne[index] < 58)) {
				encountered = false
			}
		}

	}
	return string(mutableOne)
}

func getLenCap(s string) int {
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
