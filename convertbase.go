package piscine

import (
	"github.com/01-edu/z01"
)

//ConvertBase is a function
func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := atoiBaseConvertBase(nbr, baseFrom)

	return decimalToBase(decimal, baseTo)
}

func atoiBaseConvertBase(s string, base string) int {
	if !isValidBaseForAtoiBaseConvertBase(base) {
		return 0
	}
	length := getLenOfAtoiBaseConvertBase(s)
	radix := getLenOfAtoiBaseConvertBase(base)
	baseRad := 1
	result := 0
	mutableOne := []rune(s)
	for index := length - 1; index > -1; index-- {
		digit := getIndexConvertBase(base, mutableOne[index])
		result += digit * baseRad
		baseRad *= radix
	}
	return result
}

func getLenOfAtoiBaseConvertBase(str string) int { //validated
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

func getIndexConvertBase(base string, char rune) int {
	count := 0
	for _, v := range base {
		if char == v {
			return count
		}
		count++
	}
	return count
}

func isValidBaseForAtoiBaseConvertBase(base string) bool {
	length := getLenOfAtoiBaseConvertBase(base)

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

func decimalToBase(nbr int, base string) string {
	if !isValidBaseConvertBase(base) {
		printStrForBaseConvertBase("NV")
		return ""
	}
	return convertNbrIntoBaseConvertBase(nbr, base)
}

func printStrForBaseConvertBase(str string) { //validated
	mutableStr := []rune(str)
	for s := range str {
		z01.PrintRune(mutableStr[s])
	}
}

func isValidBaseConvertBase(base string) bool {
	length := getLenOfAtoiBaseConvertBase(base)

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

func convertNbrIntoBaseConvertBase(n int, base string) string { //validated

	radix := getLenOfAtoiBaseConvertBase(base)
	requiredLength := requiredLengthToConvertConvertBase(n, radix)
	count := requiredLength - 1
	baseDigits := []rune(base)
	result := ""
	for count >= 0 {
		getReminderFrom := n / getPowerConvertBase(radix, count)
		digit := getReminderFrom % radix
		result = result + string(baseDigits[digit])
		count--
	}
	return result
}

func getPowerConvertBase(nb int, power int) int {
	result := 1
	for index := 0; index < power; index++ {
		result *= nb
	}
	return result
}

func requiredLengthToConvertConvertBase(n int, b int) int { //validated
	if n == 0 {
		return 1
	}
	length := 0
	if n > 0 {
		for n > 0 {
			length++
			n /= b
		}
	} else {
		for n < 0 {
			length++
			n /= b
		}
	}

	return length
}
